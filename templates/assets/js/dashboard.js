/* globals Chart:false, feather:false */

axios.get("http://127.0.0.1:8080/api/v1/sensors").then(
  (response) => {
      var result = response.data;
      console.log(result);
  },
  (error) => {
      console.log(error);
  }
);

(function () {
    'use strict'
  
    feather.replace({ 'aria-hidden': 'true' })
  
    // Graphs
    var ctx = document.getElementById('myChart')
    var metrics = [
      10000,
      21345,
      18483,
      24003,
      23489,
      24092,
      12034
    ]
    // eslint-disable-next-line no-unused-vars
    var myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: [
          'Sunday',
          'Monday',
          'Tuesday',
          'Wednesday',
          'Thursday',
          'Friday',
          'Saturday'
        ],
        datasets: [{
          data: metrics,
          lineTension: 0,
          backgroundColor: 'transparent',
          borderColor: '#007bff',
          borderWidth: 4,
          pointBackgroundColor: '#007bff'
        }]
      },
      options: {
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: false
            }
          }]
        },
        legend: {
          display: false
        }
      }
    })
  })()
  