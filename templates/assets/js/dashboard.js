/* globals Chart:false, feather:false */

(function () {
  'use strict'
  feather.replace({ 'aria-hidden': 'true' })
})()

var api = "http://127.0.0.1:8080/api/v1";
var localurl = "http://127.0.0.1:5500/templates";

function PlaceChart(id) {
  axios.get(api+"/monitoring/"+id).then(
    (response) => {
        var result = response.data.metrics
        var labels = [], metrics = [], len = result.length
        
        if (len > 30) {
          len = 30
        } 
        
        for(var i=0; i < len; i++) {
          metrics[i] = result[i].data
          labels[i] = result[i].dibuat_pada
        }
        
        // Graphs
        var ctx = document.getElementById('myChart')
      
        // eslint-disable-next-line no-unused-vars
        var myChart = new Chart(ctx, {
          type: 'line',
          data: {
            labels: labels,
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
  });
}

function GetSensors() {
  axios.get(api+"/sensors").then(
  (response) => {
      var ctx = document.getElementById('sensors')
      var result = response.data.sensors;
      var text = ""
      var no = 1
      result.forEach(element => {
          text += "<tr><td>"
              +no
              +"</td><td><a href=\""+localurl+"/monitoring.html?id="+element.id.Int64+"\">"
              +element.tipe
              +"</a></td><td>"
              +element.ditempatkan_pada.Time
              +"</td><td>"
              +element.provinsi+", "+ element.kecamatan +", "+ element.desa
              +"</td><td>"
              +element.status.Bool
              +"</td></tr>"
          no++
      });
      ctx.innerHTML = text;
  },
  (error) => {
      console.log(error);
  }
  );
}

function getMetrics(id) {
  axios.get(api+"/monitoring/"+id).then(
    (response) => {
        return response.data.metrics;
  });
}