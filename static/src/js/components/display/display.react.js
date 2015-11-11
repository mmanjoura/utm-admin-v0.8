/**
 * Copyright (c) 2014, U-blox.
 * All rights reserved.
 */
var React = require('react');
var AppStore = require('../../stores/app-store.js');
var Configure = require('../panels/configure.react')
var Summary = require('../panels/summary.react')
var Hint = require('../panels/hint.react')
var Measurements = require('../panels/measurements.react')
var StoreWatchMixin = require('../../mixins/StoreWatchMixin');
var DisplayRow = require('./displayRow.react')

var Link = require('react-router-component').Link

var arrData = [];

var currentUuidsObject = new Object();
var UuidsMap = new Map();
var totalMsg = 0;
var totalBytes = 0;

var Display = React.createClass({

   getInitialState: function(){   

      var data = {data:[
        ]}    

      return data;
  },
  componentDidMount: function() {
    pollState(function(data) {
      // fixup missing state properties to avoid muliple levels of missing attribute tests
      [
        "Connection",
        "LatestRssi",
        "LatestRssiDisplay",
        "LatestPowerState",
        "LatestPowerStateDisplay",
        "LatestDataVolume",
        "InitIndUlMsg",
        "LatestPollIndUlMsg",
        "LatestTrafficReportIndUlMsg",
        "LatestDisplayRow"
      ].map(function(property) {
        if (!data[property]) {
          data[property] = {};
        }
      });

      if (window.location.hash == "#debug") {
        data.json = JSON.stringify(data, null, "  ");
      } else {
        data.json = "";
      }

      var currentUuIdkey = data["LatestDisplayRow"]["Uuid"];

      currentUuidsObject[data["LatestDisplayRow"]["Uuid"]] = data;
      UuidsMap.set(data["LatestDisplayRow"]["Uuid"], data);



      if (!(currentUuIdkey in currentUuidsObject)){
        
        if(currentUuIdkey !== undefined){
              if(currentUuIdkey.length > 0){
               // data["LatestDisplayRow"]["TotalMsg"] = data["LatestDisplayRow"]["TotalMsg"] + data["LatestDisplayRow"]["UTotalMsg"] + data["LatestDisplayRow"]["DTotalMsg"] 
                console.log(data["LatestDisplayRow"]["TotalMsg"]);
                currentUuidsObject[currentUuIdkey] = data;
                UuidsMap.set(currentUuIdkey, data);
              }
        }
    
        
      }

     // arrData.push(data);
      this.setState({data: UuidsMap})

    }.bind(this), 10000);

  },

  render:function(){
    return (
            <div><br />
              <Configure />
              <Summary data = { this.state.data} />
              <Hint />
              <DisplayRow data = { this.state.data} />
            </div>
        );

        }
});


function formatTime(ts) {
  if (ts != null) {
    var i = ts.indexOf(".");
    return ts.substr(0, i).replace("T", " ");
  }

}

function pollState(updateState) {
  function pollLoop() {
    var x = new XMLHttpRequest();
    x.onreadystatechange = function() {
      if (x.readyState == 4) {
        if (x.status == 200) {
          var data = JSON.parse(x.responseText);
         
          updateState(data);
        }
        window.setTimeout(pollLoop, 1000);
      }
    };

    x.open("GET", "latestState", true);
    x.send();
  }
  pollLoop();
}

Object.size = function(obj) {
    var size = 0, key;
    for (key in obj) {
        if (obj.hasOwnProperty(key)) size++;
    }
    return size;
};

module.exports = Display;



