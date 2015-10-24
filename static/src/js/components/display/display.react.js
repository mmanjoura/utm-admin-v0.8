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
        <div className="row"><br />
          <div className="col-lg-4">
           {this.state.json}
            {/* Device Setting */}
            <div className="panel" style={{height: 60, width: 180, marginTop: 100}}>
              <div className="panel-body"><p>
                  <a tabIndex={-1} href="#/standard"> <b className="fa fa-cogs  fa-2x" /></a> <b style={{float: 'right'}}>Configure Ticked</b><br />
                </p>
              </div>
            </div>
          </div> 
          {/* /.col-lg-4 */} 
          <div className="col-lg-4">
            {/* TrxSummary */}
            <div className="panel panel-info" style={{height: 110, width: 350}}>
              <div className="panel-body">
                <p style={{fontStyle: 'italic'}}>
                  <b>Total Msg:</b> <span className="resetColor">  {this.state["LatestDisplayRow"]["TotalMsgs"]}</span><br />
                  <b>Total Bytes:</b> {this.state["LatestDataVolume"]["UplinkBytes"]}<br />
                  <b>Last Msg:</b>    {this.state["LatestDisplayRow"]["UlastMsgReceived"]}<br />
                </p> 
              </div>
            </div>
          </div>
          <div className="col-lg-4">
            {/* Hint */}
            <div className="panel panel-info" style={{height: 110, width: 300}}>
              <div className="panel-body">
                <p>
                  
                    There are more options / setting for individual UTM. They can be access through the following icon 
                   <b className="fa fa-cogs" /> in the table below.
                  
                </p>
              </div>
            </div>
          </div>
          {/* /.col-lg-12 */}
        </div>
        {/* /.row */}
        <div className="row">
          <div className="panel panel-default">
            <div className="panel-heading">
              {/* /.panel-heading */}
              <div className="panel-body">
                <div className="dataTable_wrapper">
                  <table className="table table-striped table-bordered table-hover" id="dataTables-example">
                    <thead>
                      <tr className="info">
                        <th> <input type="checkbox" style={{width: 15}} /> All</th>
                        <th>Name/Uuid</th>
                       <th>Upstream</th>
                        <th>Downsteam</th>
                        <th>RSRP </th>
                        <th>
                          Battery - <i className="fa fa-floppy-o" />
                        </th>
                      </tr>
                    </thead>
                    <tbody style={{fontSize: 12}}>
                      <tr className="even gradeA">
                        <td style={{width: 15}}>
                          <a tabIndex={-1} href="#/standardtwo"> <b className="fa fa-cogs" /></a><br />
                          <input type="checkbox" style={{width: 15}} /><br />
                          <img src="static/dist/assets/images/green.png" alt="logo" style={{maxWidth: 12}} />
                        </td>
                        <td>
                          <ul>
                            <li><b>Uuid:</b> {this.state["LatestDisplayRow"]["Uuid"]}</li>
                            <li><b>Mode:</b> {this.state["LatestDisplayRow"]["Mode"]}</li>
                            <li><b>Name:</b> {this.state["LatestDisplayRow"]["UnitName"]}</li>
                          </ul>   
                        </td>
                        <td>
                          <ul>
                            <li><b>Total Msg:</b> {this.state["LatestDisplayRow"]["TotalMsgs"]}</li>
                            <li><b>Total Bytes:</b> {this.state["LatestDataVolume"]["UplinkBytes"]}</li>
                            <li><b>Last Msg RX:</b> {this.state["LatestDisplayRow"]["UlastMsgReceived"]}</li>
                          </ul> 
                        </td>
                        <td className="center">
                          <ul>
                            <li><b>Total Msg:</b> {this.state["LatestDisplayRow"]["DTotalMsgs"] ? "" : "--"}</li>
                            <li><b>Total Bytes:</b> {this.state["LatestDisplayRow"]["DTotalBytes"] ? "" : "--"}</li>
                            <li><b>Last Msg RX:</b> {this.state["LatestDisplayRow"]["DlastMsgReceived"] ? "" : "--"}</li>
                          </ul>
                        </td>
                        <td className="center">{this.state["LatestDisplayRow"]["BatteryLevel"] ? "" : "0"}</td>
                        <td className="center">
                          {this.state["LatestDisplayRow"]["BatteryLevel"] ? "" : "0"} -
                          {this.state["LatestDisplayRow"]["DiskSpaceLeft"] ? "" : " 0"}
                        </td> 
                      </tr>
                    <tr className="even gradeC">
                        <td style={{width: 15}}>
                          <a tabIndex={-1} href="#/standardtwo"> <b className="fa fa-cogs" /></a><br />
                          <input type="checkbox" style={{width: 15}} /><br />
                          <img src= "static/dist/assets/images/green.png" alt="logo" style={{maxWidth: 12}} />
                        </td>
                        <td style={{width: 205}}>
                          <ul>
                            <li><b>Uuid:</b> {this.state["LatestDisplayRow"]["Uuid"]}</li>
                            <li><b>Mode:</b> {this.state["LatestDisplayRow"]["Mode"]}</li>
                            <li><b>Name:</b> {this.state["LatestDisplayRow"]["UnitName"]}</li>
                          </ul>   
                        </td>
                        <td style={{width: 75}}>
                          <ul>
                            <li><b>Total Msg:</b> {this.state["LatestDisplayRow"]["TotalMsgs"]}</li>
                            <li><b>Total Bytes:</b> {this.state["LatestDataVolume"]["UplinkBytes"]}</li>
                            <li><b>Last Msg RX:</b> {this.state["LatestDisplayRow"]["UlastMsgReceived"]}</li>
                          </ul> 
                        </td>
                        <td className="center" style={{width: 75}}>
                          <ul>
                            <li><b>Total Msg:</b> {this.state["LatestDisplayRow"]["DTotalMsgs"] ? "" : "--"}</li>
                            <li><b>Total Bytes:</b> {this.state["LatestDisplayRow"]["DTotalBytes"] ? "" : "--"}</li>
                            <li><b>Last Msg RX:</b> {this.state["LatestDisplayRow"]["DlastMsgReceived"] ? "" : "--"}</li>
                          </ul>
                        </td>
                        <td className="center" style={{width: 25}}>{this.state["LatestDisplayRow"]["BatteryLevel"] ? "" : "0"}</td>
                        <td className="center" style={{width: 25}}>
                          {this.state["LatestDisplayRow"]["BatteryLevel"] ? "" : "0"} -
                          {this.state["LatestDisplayRow"]["DiskSpaceLeft"] ? "" : " 0"}
                        </td> 
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              {/* /.panel-body */}
            </div>
            {/* /.panel */}
          </div>
          {/* /.col-lg-12 */}
        </div>
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



module.exports = Display;

