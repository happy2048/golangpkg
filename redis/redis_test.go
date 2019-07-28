package redis
import(
	"fmt"
	"testing"
	cv "github.com/smartystreets/goconvey/convey"
)
var (
	comment = "TF: Test Function,C: Condition,SR: ShouldReturn"
)
func TestRedis(t *testing.T) {
	cv.Convey("TestRedis",t,func(){
		fmt.Printf("\n%s\n",comment)
		cv.Convey(`{ TF: TestRedisCanConnect,C: ("127.0.0.1:6382"),SR: (nil) } ==> `,func(){
			cv.So(RedisTestConnect("127.0.0.1:6382"),cv.ShouldEqual,nil)
		})
		
	})

}
