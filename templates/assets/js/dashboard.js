/* globals Chart:false, feather:false */

(function () {
  'use strict'
  feather.replace({ 'aria-hidden': 'true' })
})()

var api = "http://127.0.0.1:8080/api/v1";
var localurl = "http://127.0.0.1:5500/templates";

function HrefValue() {
  var dashboard = document.getElementById("to_dashboard")
  dashboard.setAttribute('href', localurl+"/index.html")
}

function Export(id) {
  window.location.href = api+"/export/sensor/"+id
}

function PlaceChart(id) {
  axios.get(api+"/monitoring/"+id).then(
    (response) => {
        var result = response.data.metrics
        var labels = [], metrics = [], len = result.length
        var htmltemplate = ""
        var placeMetrics = document.getElementById('metrics')
        for (var i = 0; i < len; i++) {
          htmltemplate += "<tr><td>" 
          +(i+1)
          +"</td><td>"
          +result[i].data
          +"</td><td>"
          +result[i].dibuat_pada
          +"</td></tr>";
        }
        
        placeMetrics.innerHTML = htmltemplate
        console.log(result)
        // Graphs
        var ctx = document.getElementById('myChart'), count = len-1, end = len
        if(len > 15){
          count = 14
          end = 15
        }
        
        for(var i=0; i < end; i++) {
          metrics[i] = result[count].data
          labels[i] = result[count].dibuat_pada
          count--
        }
      
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
              +"</td></tr>";
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