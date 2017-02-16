package models


import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
    Convey("Given a Player dao", t, func() {
        p := PlayerDAO()

        Convey("When inserting players", func() {


        }

        Convey("When calling GetAll", func() {
            players := p.GetAll()

            Convey("Should return list of players", func() {

                for i, player := range players {
                  So(player, ShouldEqual, 2)
                }
            })
        })
    })
}
