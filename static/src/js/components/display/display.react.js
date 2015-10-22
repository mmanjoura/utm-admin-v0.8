/**
 * Copyright (c) 2014, U-blox.
 * All rights reserved.
 */
var React = require('react');
var Configure = require('../panels/configure.react')
var Summary = require('../panels/summary.react')
var Hint = require('../panels/hint.react')
var Measurements = require('../panels/measurements.react')
var Link = require('react-router-component').Link

var Display = React.createClass({

  getInitialState: function(){   

      var state = {

        "Connection": {},
        "LatestGpsPosition": {},
        "LatestLclPosition": {},
        "LatestLclPositionDisplay": {},
        "LatestRssi": {},
        "LatestRssiDisplay": {},
        "LatestPowerState": {},
        "LatestPowerStateDisplay": {},
        "LatestDataVolume": {},
        "LatestDisplayRow": {},
        "Device_Setting" : {}
    
    };

    if (window.location.hash == "#debug") {
      state.json = JSON.stringify(state, null, "  ");
    } else {
      state.json = "";
    }
    return state;

      return state;
  },


  componentDidMount: function() {
    pollState(function(state) {
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
        if (!state[property]) {
          state[property] = {};
        }
      });

      if (window.location.hash == "#debug") {
        state.json = JSON.stringify(state, null, "  ");
      } else {
        state.json = "";
      }
      this.setState(state);
    }.bind(this), 10000);

  },

  /**
   * @return {object}
   */
  render: function() {

    return (


      <div>
           <div>
              <Link href="/home">View Sample App</Link>
            </div>
        <div className="row"><br />
          <Configure />
          <Summary />
          <Hint />
        </div>
        <div className="row">
              <Measurements />
        </div>
      </div>
  
    );
  }
});

module.exports = Display;

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
          var state = JSON.parse(x.responseText);
         
          updateState(state);
        }
        window.setTimeout(pollLoop, 1000);
      }
    };

    x.open("GET", "latestState", true);
    x.send();
  }
  pollLoop();
}


