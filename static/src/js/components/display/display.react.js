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

      var data = {

        "Connection": {},
        "LatestGpsPosition": {},
        "LatestLclPosition": {},
        "LatestLclPositionDisplay": {},
        "LatestRssi": {},
        "LatestRssiDisplay": {},
        "LatestPowerState": {},
        "LatestPowerStateDisplay": {},
        "LatestDataVolume": {},
        "LatestDisplayRow": {}
    
    
    };

    if (window.location.hash == "#debug") {
      data.json = JSON.stringify(data, null, "  ");
    } else {
      data.json = "";
    }
    return data;

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
      this.setState(data);
    }.bind(this), 10000);
  },

  render:function(){

    return (
      <div className="row">
      <div>
        <div className="row"><br />
         <Configure />
          {/* /.col-lg-4 */} 
          <Hint />
          <div className="col-lg-4">
            {/* TrxSummary */}
            <div className="panel panel-info" style={{height: 110, width: 400}}>
              <div className="panel-body">
                <p style={{fontStyle: 'italic'}}>
                  <b>Total Msg:</b> <span className="resetColor">  {this.state["LatestDisplayRow"]["TotalMsgs"]}</span><br />
                  <b>Total Bytes:</b><span className="resetColor">  {this.state["LatestDataVolume"]["UplinkBytes"]}</span><br />
                  <b>Last Msg:</b> <span className="resetColor">    {this.state["LatestDisplayRow"]["UlastMsgReceived"]}</span><br />
                </p> 
              </div>
            </div>
          </div>
        
  
        </div>
        {/* /.row */}
        <div className="row" >
          <div className="panel panel-default" >
            <div className="_panel-heading" style={{width:'100%'}}>
              {/* /.panel-heading */}
              <div className="panel-body">
                <div className="dataTable_wrapper">
                  <table className="table table-striped table-bordered table-hover" id="dataTables-example">
                    <thead>
                      <tr className="info">
                        <th> <input type="checkbox" style={{width: 15}} /> All</th>
                        <th>Name/Uuid</th>
                       <th>Uplink</th>
                        <th>Downlink</th>
                        <th>RSRP </th>
                        <th>
                          Others
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
                          <ul className="SmallPadding">
                            <li ><b>Uuid:</b>{this.state["LatestDisplayRow"]["Uuid"]}</li>
                            <li><b>Mode:</b> {this.state["LatestDisplayRow"]["Mode"]}</li>
                            <li><b>Name:</b> {this.state["LatestDisplayRow"]["UnitName"]}</li>
                             <li><b>Reporting Interval:</b> {this.state["LatestDisplayRow"]["ReportingInterval"]}</li>
                              <li><b>Heart Beat:</b> {this.state["LatestDisplayRow"]["HeartbeatSeconds"]}</li>
                          </ul>   
                        </td>
                        <td>
                          <ul className="SmallPadding">
                            <li><b>Total Msg:</b> {this.state["LatestDisplayRow"]["TotalMsgs"]}</li>
                            <li><b>Total Bytes:</b> {this.state["LatestDataVolume"]["UplinkBytes"]}</li>
                            <li><b>Last Msg RX:</b> {this.state["LatestDisplayRow"]["UlastMsgReceived"]}</li>
                          </ul> 
                        </td>
                        <td className="center">
                          <ul className="SmallPadding">
                            <li><b>Total Msg:</b> {this.state["LatestDisplayRow"]["DtotalMsgs"]}</li>
                            <li><b>Total Bytes:</b> {this.state["LatestDisplayRow"]["DTotalBytes"]}</li>
                            <li><b>Last Msg RX:</b> {this.state["LatestDisplayRow"]["DlastMsgReceived"]}</li>
                          </ul>
                        </td>
                        <td className="center">{this.state["LatestDisplayRow"]["RSRP"]}</td>
                        <td className="center" style={{width: 105}}>
                          <i className="fa fa-floppy-o" />&nbsp; { this.state["LatestDisplayRow"]["DiskSpaceLeft"]}<br />
                          <i className="fa fa-battery-full" />&nbsp; { this.state["LatestDisplayRow"]["BatteryLevel"]}
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
      </div>
    )
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

module.exports = Display;



