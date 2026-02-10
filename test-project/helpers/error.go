package helpers

import "fmt"

type Errotype struct{}

func (Errotype) GetData(data interface{}, err error) string {
	return fmt.Sprintf("error get data %v : %v", data, err)
}

func (Errotype) SaveData(data interface{}, err error) string {
	return fmt.Sprintf("error save data %v : %v", data, err)
}

func (Errotype) CreateData(data interface{}, err error) string {
	return fmt.Sprintf("error create data %v : %v", data, err)
}

func (Errotype) PushData(data interface{}, err error) string {
	return fmt.Sprintf("error push data %v : %v", data, err)
}

func (Errotype) UpdateData(data interface{}, err error) string {
	return fmt.Sprintf("error update data %v : %v", data, err)
}

func (Errotype) DeleteData(data interface{}, err error) string {
	return fmt.Sprintf("error delete data %v : %v", data, err)
}

func (Errotype) Marshal(data interface{}, err error) string {
	return fmt.Sprintf("error marshal %v : %v", data, err)
}

func (Errotype) UnMarshal(data interface{}, err error) string {
	return fmt.Sprintf("error unmarshal %v : %v", data, err)
}

func (Errotype) BindJSON(data interface{}, err error) string {
	return fmt.Sprintf("error bind json %v : %v", data, err)
}

func (Errotype) SendData(data interface{}, err error) string {
	return fmt.Sprintf("error send data %v : %v", data, err)
}

func (Errotype) Decode(data interface{}, err error) string {
	return fmt.Sprintf("error decoding %v : %v", data, err)
}

func (Errotype) Encode(data interface{}, err error) string {
	return fmt.Sprintf("error encoding %v : %v", data, err)
}

func (Errotype) Parse(data interface{}, err error) string {
	return fmt.Sprintf("error parsing data %v : %v", data, err)
}

func (Errotype) Write(data interface{}, err error) string {
	return fmt.Sprintf("error write data %v : %v", data, err)
}

func (Errotype) Read(data interface{}, err error) string {
	return fmt.Sprintf("error read data %v : %v", data, err)
}
