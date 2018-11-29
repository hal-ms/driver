package service

//
//import (
//	"github.com/tarm/serial"
//)
//
//var Building = newBuildingService()
//
//type buildingService struct {
//	conn *serial.Port
//}
//
////
////func newBuildingService() buildingService {
////	c := &serial.Config{Name: "COM14", Baud: 9600}
////	s, err := serial.OpenPort(c)
////	if err != nil {
////		//log.SendSlack(err.Error())
////		panic(err)
////	}
////	conn := s
////	return buildingService{conn: conn}
////}
////
////func (s *buildingService) Animation(scene string) error {
////	db := db.DB.Get()
////	var d model.Scene
////	err := db.Read("Scene", scene, &d)
////	if err != nil {
////		return errors.New("存在しないシーン")
////	}
////	s.write(d.Raw)
////	return nil
////}
////
////// 信号送信
////func (s *buildingService) write(b []byte) error {
////	_, err := s.conn.Write(b)
////	if err != nil {
////		//log.SendSlack(err.Error())
////		panic(err)
////	}
////	return nil
////}
