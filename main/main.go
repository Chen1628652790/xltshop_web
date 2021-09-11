package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

func main() {
	appID := "2016102500758519"
	privateKey := "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCQTTrd/hUWDM0EOaCbknmzLTzLE8BTZO9oQ0n8PiZhXf4gsn3W6wfhuZ+TzvC1yJUl9g1eysnPMGTe5Iz0TO1ZaVnhVd89fHYpJGcr4XhGuCSNEEwqaKHa/3GfK9LzFluVKR4b/CaRsRvLrRy1nVzX+1wWR7fBalpycwsloaEg0Ps9N90Ih/iogdcs+Ex3W5rnW/mAak3rJaMFrVWpSeJzf9aTv/GTFjB0FwOnRM3ILV4WS4dMVKZS+U09V9ds+1e+unSB94GmiD9MBpu5FZaBIJjxvasIpRIguuQpixDCtRDT1Vkpa8T9iN9pS0StypOd7cMiHGvpAE8NorOXZbahAgMBAAECggEAPyPL0j7O8yr+ug6NHAVnguMUardlvBe6OaDXyqtXF9uMyrnPHi8Q78/M51vxL1lpCYc0KnoI+8NtH6pZkrvmTu0uCs8MM1c2TKJFEopBmpAQTjkHWrcVu0Fycfc3Am6R/B5VsmEOb0lTpDdHDKCic07k6ErROKxUjyePhRPH8RdSZpNJPKd/uTUdZVOc8JGjpGAdmNT9q7KrLpteMuYzMEcoLXlhitGkC2ELYtg/MSL8LJ94UPqe8HAPPnifBE0ShZ3anDf62QLwHRuWTJA1ArMaDiZuNSu5Gk+gXk3PINDXPwyKUneBWNYsGHHB1EgwQiKDpsJmuy9WysCpydGewQKBgQD3/Pui3Rrh9Io9CwW0x/wiX/4sBXWImJcGTIiLX8biURSfFVJEAtznxPQf7+BuiW+lDn7ieT8pBKl8zACs+GrIDXf7AEuLUQwRF2eHuoAUXqtO4fQlO9EAe/PjhzYnTyJZeqEDc2qWdPUosRcCOAj6rkO/wdogrSqU1ww+yBErhQKBgQCU9rHa+Br0M5WzG3mpHAyNSymmQ1Vp3WROgioRZlx8vOBW10uDUUaKzwPVXLq8E7EkMpUSYG5Mm1h+kw0gjj0npa++OnreyurDHOnAZcwYBCz4oh6dXy8ADnhrB9Vz33yltYPPvQ8J5IhitTlGOhqLYbHjlVcigs519D6kbIcjbQKBgQDg7NKH66egXh6sMz0ftWvY+dwdrW3nUQ9aJTyLvXk0eHoNuFb/XOFkTl1mQjn7yCg9OyKW44YH/DSF/rp2KHMhtCWowaHDYOVi8ylyEBRvZVZXm9XHl7N/ju2s50yU5s+u/OzhpmN8x+Q83jKSTqTGSh0k5fykOqwuh8aRgwEfyQKBgQCHNuZ6PHs67xgTW/0y46MBHhjQMo51aeCC4uQMpz2MfGWmbga9TCkcFo0EPwfBcJ56nO9zntR/9QJ+4jwoMPSR6HN92NdvJAG02anUWpLHugKYLZBciOnAw2HKxXGbnGEoiXr8NkBQWoDyGE3E0TkHC8bNLeHKEbIWn329AkYogQKBgQDoPSy5diUILU3tYK9o/slix+WhxuZN09dbrGd2Dzm6NZLRuJzjpL0QXURHq3tbXbgdynLG4XTKdbDCLcxjxAzByDBLX6v6p0Vx90elhvvOwg2/R+QLNFILC4AfkOF7qGEeqaQ3NgsLDQYlqClbs4YOf2XFcKi0xDJ9Om72ccsA6w=="
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoqrs44zipW5XFHsyz9TxLSe0jmF5JKhatut/xsHSxyAgrio27mavyX8SmSoqtD3Fz+/wKHZoAtH63guCju62SseMukG2N3dLAv1h2kGoZXmeqJLTU1BBnJTdfrNeaEWZaed1TXgO2ZaTdd16NOGgHwz6U6QRNuec1D1FQxdYdHAEZB0h5xvIzaqRYQGLEL9kStea9jDd8fNv4CRAxcp7H5isTFmm96VGqHfAV1qXyd7shKtYuPKBuQPwU+GQhf5lkXTFhNg0dK1kV6oMAaDlg6K3n4cmWR6mEZFhhXizYJrqYJgs15Y52q1iQ0IdctxmCl5UZrIVaQ38uxJcz3QK6QIDAQAB"

	client, err := alipay.New(appID, privateKey, false)
	if err != nil {
		panic(err)
	}

	if err := client.LoadAliPayPublicKey(aliPublicKey); err != nil {
		panic(err)
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://orders.free.svipss.top/o/v1/pay/alipay/notify"
	p.ReturnURL = "http://edu.xiaolatiao.cn"
	p.Subject = "生鲜商城-订单支付"
	p.OutTradeNo = "1628652790"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}
	var payURL = url.String()
	fmt.Println(payURL)
}
