package main

import (
	"log"
	"os/exec"
	"strings"
	"io/ioutil"
)



func main() {

	out, err := exec.Command("sh", "nmap-run.sh").Output();

	if err != nil {

		log.Fatal(err);
	}


	macs := strings.Replace(string(out), "MAC Address: ", "", -1);
	macs = strings.Replace(macs, " (Unknown)", ",", -1);

	macs = macs[0:strings.LastIndex(macs, ",")];

	macsArr := strings.Split(macs, ",");

	content, err := ioutil.ReadFile("known-mac-addr");

	if err != nil {

		log.Fatal(err);
	}

	kmacArr := strings.Split(string(content), "\n");

	isMacFound := false;

	var uknownMacs string = "";

	for i := 0; i < len(macsArr); i++ {

		for j := 0; j < len(kmacArr); j++ {

			if macsArr[i] == kmacArr[j] {

				isMacFound = true;
				break;
			}

		}

		if isMacFound == false {
			uknownMacs += macsArr[i] + ",";
		}
	}



	if isMacFound == false {

		uknownMacs = macs[0:strings.LastIndex(uknownMacs, ",")];
		err := exec.Command("sh", "email-alert.sh", uknownMacs).Run();

		if err != nil {

			log.Fatal(err);
		}

	}

}
