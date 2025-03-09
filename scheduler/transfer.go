package scheduler

import (
	"context"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"pool/dao"
	"pool/log"
	"pool/models"
	"sync"
	"time"
)

func NewHTTPTransport(timeout time.Duration, maxIdleConnsPerHost int, keepAlive time.Duration) *http.Transport {
	proxyURL, _ := url.Parse("http://127.0.0.1:7890")
	return &http.Transport{
		IdleConnTimeout:     timeout,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		Proxy:               http.ProxyURL(proxyURL),
		Dial: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
		}).Dial,
	}
}

func NewHTTP(timeout time.Duration, maxIdleConnsPerHost int, keepAlive time.Duration) *http.Client {
	transport := NewHTTPTransport(timeout, maxIdleConnsPerHost, keepAlive)
	return &http.Client{Timeout: timeout, Transport: transport}
}

func NewRPC(rpcEndpoint string) *rpc.Client {
	var (
		defaultMaxIdleConnsPerHost = 10
		defaultTimeout             = 25 * time.Second
		defaultKeepAlive           = 180 * time.Second
	)
	opts := &jsonrpc.RPCClientOpts{
		HTTPClient: NewHTTP(
			defaultTimeout,
			defaultMaxIdleConnsPerHost,
			defaultKeepAlive,
		),
	}

	rpcClient := jsonrpc.NewClientWithOpts(rpcEndpoint, opts)

	return rpc.NewWithCustomRPCClient(rpcClient)
}

func decodeSystemTransfer(tx *solana.Transaction) error {
	for _, instr := range tx.Message.Instructions {
		if instr.ProgramIDIndex == 2 {
			accounts, err := instr.ResolveInstructionAccounts(&tx.Message)
			if err != nil {
				return err
			}

			inst, err := system.DecodeInstruction(accounts, instr.Data)
			if err != nil {
				return err
			}

			if transferInst, ok := inst.Impl.(*system.Transfer); ok {
				lamports := transferInst.Lamports
				from := transferInst.Get(0)
				to := transferInst.Get(1)
				lamportsOnAccount := new(big.Float).SetUint64(*lamports)
				solBalance := new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

				fmt.Println("from: ", from)
				fmt.Println("to: ", to)
				fmt.Println("◎", solBalance.Text('f', 10))

				// 保存到数据库
				// todo: 判断接收地址是否正确
				// 查询地址的用户
				fromUserId, err := dao.GetUserIdByAddress(from.PublicKey.String())
				if err != nil {
					return err
				}

				// 保存到数据库
				transfer := &models.Transfer{
					UserId: fromUserId,
					Amount: *solBalance,
				}
				err = dao.AddTransfer(transfer)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func MonitorTransfer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		log.SystemLog().Warnf("MonitorTransfer ws connect err:%v", err)
		return
	}

	program := solana.MustPublicKeyFromBase58("2XN2N3973SaswCRFdcrYS7GGuDnFvjeAZCVLtJddQusW")

	subscription, err := wsClient.LogsSubscribeMentions(program, rpc.CommitmentRecent)
	if err != nil {
		log.SystemLog().Warnf("MonitorTransfer err:%v", err)
		return
	}
	defer subscription.Unsubscribe()

	for {
		recv, err := subscription.Recv(ctx)
		if err != nil {
			log.SystemLog().Warnf("monitor transfer err:%v", err)
			continue
		}

		// 查询交易信息
		client := NewRPC(rpc.DevNet_RPC)
		transaction, err := client.GetTransaction(
			ctx,
			recv.Value.Signature,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingBase64,
			},
		)
		if err != nil {
			log.SystemLog().Warnf("monitor transfer err:%v", err)
			continue
		}

		decodedTx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(transaction.Transaction.GetBinary()))
		if err != nil {
			log.SystemLog().Warnf("monitor transfer err:%v", err)
			continue
		}

		err = decodeSystemTransfer(decodedTx)
		if err != nil {
			log.SystemLog().Warnf("monitor transfer err:%v", err)
			continue
		}
	}
}
