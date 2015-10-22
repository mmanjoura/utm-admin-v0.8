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

function getUtmsData(){
  return { utmsData: AppStore.getUtmsData()}
}

var Display = React.createClass({

  getInitialState: function(){   
    return getUtmsData();
  },

  render: function() {
    var utmsData = this.state.utmData.map(function(utmData){
      <tr>
        <td>utmData["Uuid"]</td>
        <td>utmData["Mode"]</td>
        <td>utmData["UnitName"]</td>
        <td>utmData["UTotalMsgs"]</td>
        <td>utmData["UTotalMsgs"]</td>
         <td>utmData["UTotalMsgs"]</td>d>
      </tr>
    })
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


