package ethutils

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/chain/eth/utils/hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
	"testing"
)

func TestGenerateRandomAddress(t *testing.T) {

	address, pk := GenerateRandomAddress()
	println(address, pk)
}

//step glue moment
//tennis alter ranch
//auto valley foil
//vital rabbit parade

func TestGenerateByMnemonic(t *testing.T) {

	//entropy, err := bip39.NewEntropy(128)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	////var mnemonic = "pepper hair process town say voyage exhibit over carry property follow define"
	//fmt.Println("mnemonic:", mnemonic)
	//mnemonic = "jeans bitter notice collect exchange critic uncover vintage october grab dust meat"
	mnemonic := "落 霞 与 孤 鹜 齐 飞 秋 水 共 长 天 一 色"

	seed := bip39.NewSeed(mnemonic, "") //这里可以选择传入指定密码或者空字符串，不同密码生成的助记词不同

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1000") //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKey, _ := wallet.PublicKeyHex(account)

	fmt.Println("address0:", address)      // id为0的钱包地址
	fmt.Println("privateKey:", privateKey) // 私钥
	fmt.Println("publicKey:", publicKey)   // 公钥

	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	//account1, err := wallet.Derive(path, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//address1 := account1.Address.Hex()
	//privateKey1, _ := wallet.PrivateKeyHex(account1)
	//
	//fmt.Println("address1:", address1, privateKey1)
}
