# goranger

**Goranger** is a console application for getting IP ranges from [https://suip.biz](https://suip.biz) web-services by city, country (very big size!! i'm really afraid) or ISP.

    ███ ███ ███ ███ ███ ███ ███ ███ 
    █ █ █ █ █     █ █ █ █ █ ███ █   
    █ █ █ █ █   ███ █ █ █ █ █   █   
    ███ ███ █   ███ █ █ ███ ███ █   
      █                   █         
    ███                 ███         

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/goranger)](https://goreportcard.com/report/github.com/hIMEI29A/goranger) [![GoDoc](https://godoc.org/github.com/hIMEI29A/goranger/libgoranger?status.svg)](http://godoc.org/github.com/hIMEI29A/goranger/libgoranger) [![Apache-2.0 License](https://img.shields.io/badge/license-Apache--2.0-red.svg)](LICENSE)

## TOC
- [About](#about)
- [Install](#install)
- [Usage](#usage)

## About

**Goranger** is a simpliest (two files with code only) console application for getting IP ranges from [https://suip.biz](https://suip.biz) web-services by city, country (very big size!! i'm really afraid) or ISP. In case of ISP you may specify single ip or web-site url of that provider as argument of script. **Goranger** may be installed as application or imported into other Golang packages.

It is Golang port of deprecated PHP script [getIpByIsp](https://github.com/hIMEI29A/getIpbyIsp).

## Install

##### As application

With `dpkg`

    wget https://github.com/hIMEI29A/goranger/releases/download/0.1.1/\
    goranger-0.1.1-amd64.deb && sudo dpkg -i goranger-0.1.1-amd64.deb

With `gdebi`

    wget https://github.com/hIMEI29A/goranger/releases/download/0.1.1/\
    goranger-0.1.1-amd64.deb && sudo gdebi goranger-0.1.1-amd64.deb

Check the [release page](https://github.com/hIMEI29A/goranger/releases)!

##### As package

```sh
go get github.com/hIMEI29A/goranger/libgoranger
```

## Usage

To get help

```shell
./goranger -h
```

To get all IP ranges of **Serbia**

```shell
./goranger -t country -r RS
```

To get all IP ranges of **London** and save it to **file.txt**

```shell
./goranger -t city -r london -o file.txt
```

To get all IP ranges of **Beeline** ISP

```shell
./goranger -t isp -r 217.118.85.19
```

To get all IP ranges of ISP of **kremlin.ru**

```shell
./goranger -t isp -r kremlin.ru
```

**Warning!!!** In case of request by city, write the name of the city **carefully and accurately**, as much as possible. If an error occurs in the name, the search on the uncleaned database is activated, and the result includes **ALL** IP ranges from **all** possible variants. To get the most accurate result, the city name **must not contain errors**.
