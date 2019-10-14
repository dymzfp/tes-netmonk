package function

import (
	"github.com/dymzfp/tes-netmonk/model"
)

func Join(snmp []model.Snmp, snmptrap []model.Snmp) []model.Snmp {
	var newSnmp []model.Snmp
	d := 0

	for i, snmpValue := range snmp {
		for _, snmptrapValue := range snmptrap {
			if snmpValue.Ip == snmptrapValue.Ip && snmpValue.Interface == snmptrapValue.Interface {
				lenS := len(snmpValue.Link_status)
				lenT := len(snmptrapValue.Link_status)

				newLink := []model.Link_status{}

				for {
					if snmpValue.Link_status[0].Time < snmptrapValue.Link_status[0].Time {
						newLink = append(newLink, snmpValue.Link_status[0])
						snmpValue.Link_status = snmpValue.Link_status[1:]
						lenS -= 1
					} else {
						newLink = append(newLink, snmptrapValue.Link_status[0])
						snmptrapValue.Link_status = snmptrapValue.Link_status[1:]
						lenT -= 1
					}

					if lenS <= 0 {
						newLink = append(newLink, snmptrapValue.Link_status...)
						break
					} else if lenT <= 0 {
						newLink = append(newLink, snmpValue.Link_status...)
						break
					}
				}

				newSnmp = append(newSnmp, model.Snmp{snmp[i].ID, snmp[i].Ip, snmp[i].Interface, newLink})
			} else {
				d += 1
				if d == len(snmptrap) {
					newSnmp = append(newSnmp, model.Snmp{snmp[i].ID, snmp[i].Ip, snmp[i].Interface, snmp[i].Link_status})
				}
			}
		}

		d = 0

	}

	return newSnmp
}

func Avege(snmp []model.Snmp) []model.Avg {
	var avg []model.Avg 
	totalAktif := 0
	var rata float32
	for _, sn := range snmp {
		for _, link := range sn.Link_status {
			if link.Status == 1 {
				totalAktif += 1
			}
		}

		rata = float32(totalAktif) / float32(len(sn.Link_status)) * 100

		avg = append(avg, model.Avg{sn.Ip, sn.Interface, rata})
	
		totalAktif = 0
	}

	return avg
}

func DownTime(snmp []model.Snmp) []model.Down {
	var down []model.Down

	var downtime = 0
	var countdown = 0
	var mttr float32
	
	for _, sn := range snmp {
		for i := 0; i < len(sn.Link_status); i++ {
			if i == 0 {
				if sn.Link_status[i].Status == 2 {
					downtime += sn.Link_status[i].Time
				}

				if sn.Link_status[i].Status == 2 {
					countdown += 1
				}
			} else {

				if sn.Link_status[i].Status == 2 {
					downtime += sn.Link_status[i+1].Time - sn.Link_status[i].Time
				}

				if sn.Link_status[i].Status == 2 && sn.Link_status[i-1].Status == 1 {
					countdown += 1
				}

			}
		}

		if (downtime == 0 && countdown == 0) {
			mttr = 0.0
		} else {
			mttr = float32(downtime) / float32(countdown)
		}
		
		down = append(down, model.Down{sn.Ip, sn.Interface, downtime, countdown, mttr})

		downtime = 0
		countdown = 0
	}

	return down
}