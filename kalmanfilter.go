// https://en.m.wikipedia.org/wiki/Kalman_filter

package main

import (
  "fmt"
  // "time"
  // "math/rand"
)

/*

FilterData struct, initialize this struct before commencing 
any operations, as sensors are read, this struct must be 
updated alongside

*/
type FilterData struct {

/*
State the state sensor value. In a IMU this would be the 
Accelerometer
*/ 
  State float64

/*
Bias: the delta sensor error. This is the deviation 
from sensor reading and actual value. Bias can be caused by 
electromagnetic interference and represents a permanent error
in delta sensor reading. Bias is detected by averaging the 
delta sensor reading at stationary state of delta sensor
*/
  Bias float64

/*
Covariance Matrix a 2d 2x2 matrix (also known as dispersion 
matrix or variance-covariance matrix) is a matrix whose 
element in the i, j position is the covariance between the i
and j elements of a random vector. Leave this at default 
value of [[0,0],[0,0]]
*/
  Covariance [2][2]float64

  QAngle float64
  QBias float64
  RMeasure float64
}

/*
Call this method to update the state value based on sensor fusion of state and delta sensor and the previously calculated reading to get progressively more accurate state values
*/
func (filterData *FilterData) Update(stateReading, deltaReading, deltaTime float64 ) float64{
  rate := deltaReading - filterData.Bias
  state := filterData.State + (rate * deltaTime)

  filterData.Covariance[0][0] += deltaTime * (deltaTime * filterData.Covariance[1][1] - filterData.Covariance[0][1] - filterData.Covariance[1][0] + filterData.QAngle)
  filterData.Covariance[0][1] -= deltaTime * filterData.Covariance[1][1]
  filterData.Covariance[1][0] -= deltaTime * filterData.Covariance[1][1]
  filterData.Covariance[1][1] += deltaTime * filterData.QBias

  innovationCovariance := filterData.Covariance[0][0] + filterData.RMeasure

  kalmanGain := []float64{filterData.Covariance[0][0]/innovationCovariance, filterData.Covariance[1][0]/innovationCovariance}

  y := stateReading - state
  filterData.State += kalmanGain[0] * y
  filterData.Bias += kalmanGain[1] * y

  filterData.Covariance[0][0] -= kalmanGain[0] * filterData.Covariance[0][0]
  filterData.Covariance[0][1] -= kalmanGain[0] * filterData.Covariance[0][1]
  filterData.Covariance[1][0] -= kalmanGain[1] * filterData.Covariance[1][0]
  filterData.Covariance[1][1] -= kalmanGain[1] * filterData.Covariance[1][1]

  return filterData.State
}

/**********************************************************************************************************************/


func main() {
  myFilterData := FilterData {
    State     : 1,
    Bias      : 0,
    QAngle    : 1,
    QBias     : 1,
    RMeasure  : 1,
  }

  tmp := 0.0;
  duration := 1.0;
  for i := 0.0; i < 20; i++ {
    fmt.Println(tmp, i)
    // newState := myFilterData.Update(i, tmp, duration)
    newState := myFilterData.Update(tmp, tmp-i, duration)
    tmp = newState 
  }


  // var oldTime time.Time = time.Now()
  // for {
  // //   // stateReading := float64(getStateSensorReading()) // in units X
  // //   // deltaReading := float64(getDeltaSensorReading()) // in unit X per nanosecond
  //   stateReading := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
  //   deltaReading := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()

  //   var newTime time.Time = time.Now()
  //   var duration = newTime.Sub(oldTime)
  //   oldTime = newTime
  //   newState := myFilterData.Update(stateReading, deltaReading, float64(duration / time.Nanosecond))
  //   // fmt.Println(stateReading, deltaReading, float64(duration/time.Nanosecond))
  //   fmt.Println(myFilterData, newState)
  //   // fmt.Println(stateReading, deltaReading, newState)
  // }
}








