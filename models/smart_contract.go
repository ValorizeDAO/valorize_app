package models

type SmartContract struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Key      string `json:"key" gorm:"unique;"`
	ByteCode string `json:"byte_code"`
}

func (m *Model) GetSmartContractByKey(key string) (SmartContract, error) {
	var contractInfo SmartContract
	err := m.db.Where("`key` = ?", key).First(&contractInfo).Error
	if err != nil {
		return SmartContract{}, err
	}
	return contractInfo, err
}
