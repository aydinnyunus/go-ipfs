package coreunix

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"io/ioutil"
	"os"
)

func Encrypt(filename string) string {

	/* TODO: this should be replaced by input */
	content, err := ioutil.ReadFile("/home/socketpuppets/github/ipfs/public.pgp")

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	// Encrypt data with password
	armor, err := helper.EncryptMessageWithPassword(content, string(file))

	if err != nil {
		log.Fatal(err)
	}

	/* TODO: this should be replaced by input */

	err = os.WriteFile("/home/socketpuppets/github/ipfs/message.pgp", []byte(armor), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return armor

}
