package setup

import (
    "log"
    "github.com/golang/protobuf/proto"

    "github.com/machinule/nucrom/proto/gen"
)

func Try() {
    test := &pb.ProvinceSettings{
        Id: pb.ProvinceId_SWEDEN.Enum(),
        Label: proto.String("Ayy"),
    }
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    newTest := &pb.ProvinceSettings{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    // Now test and newTest contain the same data.
    if test.GetId() != newTest.GetId() {
        log.Fatalf("data mismatch %q != %q", test.GetId(), newTest.GetId())
    }
    log.Printf("Unmarshalled to: %+v", newTest)
}
