/**
 * Copyright (c) 2014, U-blox.
 * All rights reserved.
 */
var React = require('react');

var Mode = React.createClass({

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

      <div className="row"><br />
        <div className="col-lg-4">
          {/* Device Setting */}
          <div className="panel panel-info" style={{height: 100}}>
            <div className="panel-body">
              <p>
                <b>Add or Remove more UEs</b>
              </p><div>
                {/* Note the missing multiple attribute! */}
                <select id="example-multiple-selected" multiple="multiple">
                  <option value={1}>UE-P</option>
                  <option value={2} selected="selected">UE-TW</option>
                  {/* Option 3 will be selected in advance ... */}
                  <option value={3}>UE-WW</option>
                  <option value={4}>UE-T</option>
                  <option value={5}>UE-R</option>
                  <option value={6}>UE-E</option>
                </select>
              </div>
              <p />
            </div>
          </div>
        </div> 
        <div className="row">
          <div className="col-lg-12">
            <div className="panel panel-info">
              <div className="panel-heading">
                Operating Modes
              </div>
              {/* /.panel-heading */}
              <div className="panel-body">
                {/* Nav tabs */}
                <ul className="nav nav-tabs">
                  <li className="active"><a href="#standard" data-toggle="tab">Standard</a>
                  </li>
                  <li><a href="#commissioning" data-toggle="tab">Commissioning</a>
                  </li>
                  <li><a href="#traffictest" data-toggle="tab">Traffic Test</a>
                  </li>
                  <li><a href="#standalonebasic" data-toggle="tab">Standalone TRX Basic</a>
                  </li>
                </ul>
                {/* Tab panes */}
                <div className="tab-content">
                  <div className="tab-pane fade in active" id="standard">
                    <h4 style={{marginLeft: 20}}>Basic Mode</h4>
                    <p>
                    </p><div className="col-lg-12">
                      <div className="panel panel-info" style={{height: 250}}>
                        <i className="fa fa-spinner fa-spin" style={{float: 'right', padding: 10, color: 'red'}} />
                        <div className="panel-body"><p>
                          </p><div className="col-lg-8">
                            <div className="col-1-3">
                              <div className="cell red">Status: Disconnected
                              </div>
                            </div>
                          </div>
                          <div className="col-lg-8">
                            <div className="col-1-3">
                              <div className="cell grey">WAKE UP CODE: -
                              </div>
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">REVISION LEVEL: -
                              </div>
                            </div>
                            <div className="col-1-3 push-1-3">
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">LAST SEEN: - 04/10/2015
                              </div>
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">REPORTING INTERVAL: 
                              </div>
                            </div>
                            <div className="col-1-3 push-1-3">
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">HEART BEAT: -
                              </div>
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">DATE TIME: 
                              </div>
                            </div>
                          </div>
                          <p />
                        </div>
                      </div>
                    </div>
                    <p />
                  </div>
                  <div className="tab-pane fade" id="commissioning">
                    <h4 style={{marginLeft: 20}}>Commissioning Mode</h4>  
                    <p>
                    </p><div className="col-lg-12">
                      <div className="panel panel-info" style={{height: 250}}>
                        <i className="fa fa-spinner fa-spin" style={{float: 'right', padding: 10, color: 'red'}} />
                        <div className="panel-body"><p>
                          </p><div className="col-lg-8">
                            <div className="col-1-3">
                              <div className="cell red">Status: Disconnected
                              </div>
                            </div>
                          </div>
                          <div className="col-lg-8">
                            <div className="col-1-3">
                              <div className="cell grey">RSSI: -
                              </div>
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">RSRP: -
                              </div>
                            </div>
                            <div className="col-1-3 push-1-3">
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">CELL ID: - 
                              </div>
                            </div>
                            <div className="col-1-3">
                              <div className="cell grey">SNR: 
                              </div>
                            </div>
                          </div>
                          <p />
                        </div>
                      </div>
                    </div>
                    <p />
                  </div>
                  <div className="tab-pane fade" id="traffictest">
                    <h4 style={{marginLeft: 20}}>Traffic Test Mode</h4>
                    <p>
                    </p><div className="col-lg-4">
                      <div className="panel panel-info" style={{height: 300}}>
                        <div className="panel-body"><p>
                            <input className="form-control" placeholder=" No Of UL Datagrams" style={{width: 180, height: 30, margin: 5}} />
                            <input className="form-control" placeholder=" Size Of UL Datagrams" style={{width: 180, height: 30, margin: 5}} />
                            <input className="form-control" placeholder=" No Of DL Datagrams" style={{width: 180, height: 30, margin: 5}} />
                            <input className="form-control" placeholder=" Size Of DL Datagrams" style={{width: 180, height: 30, margin: 5}} /><br />
                            <button type="button" className="btn btn-info" style={{width: 100, height: 30, marginTop: 10}}>Set up Test</button>
                            <button type="button" className="btn btn-default" style={{width: 100, height: 30, marginTop: 10}}>Start Test</button>
                          </p>
                        </div>
                      </div>
                    </div>
                    <div className="panel panel-info" style={{height: 300}}>
                      <i className="fa fa-spinner fa-spin" style={{float: 'right', padding: 10, color: 'red'}} />
                      <div className="panel-body"><p>
                        </p><div className="col-lg-8" style={{float: 'right'}}>
                          <div className="col-1-3">
                            <div className="cell red">Status: Disconnected
                            </div> 
                          </div>
                        </div>
                        <div className="col-lg-8" style={{float: 'right'}}>
                          <div className="col-1-3">
                            <div className="cell grey">No of UL Datagrams: -
                            </div>
                          </div>
                          <div className="col-1-3">
                            <div className="cell grey">No of UL Bytes: -
                            </div>
                          </div>
                          <div className="col-1-3 push-1-3">
                          </div>
                          <div className="col-1-3">
                            <div className="cell grey">No of DL Datagrams: - 
                            </div>
                          </div>
                          <div className="col-1-3">
                            <div className="cell grey">No of DL Bytes: 
                            </div>
                          </div>
                          <div className="col-1-3 push-1-3">
                          </div>
                          <div className="col-1-3">
                            <div className="cell grey">No of DL Datagrams Missed: -
                            </div>
                          </div>
                        </div>
                        <p />
                      </div>
                      <div className="tab-pane fade" id="standalonebasic">
                        <h4 style={{marginLeft: 20}}>Basic Mode</h4>
                        <p>
                        </p><div className="col-lg-12">
                          <div className="panel panel-info" style={{height: 250}}>
                            <i className="fa fa-spinner fa-spin" style={{float: 'right', padding: 10, color: 'red'}} />
                            <div className="panel-body"><p>
                              </p><div className="col-lg-8">
                                <div className="col-1-3">
                                  <div className="cell red">Status: Disconnected
                                  </div>
                                </div>
                              </div>
                              <div className="col-lg-8">
                                <div className="col-1-3">
                                  <div className="cell grey">No of UL Datagrams: -
                                  </div>
                                </div>
                                <div className="col-1-3">
                                  <div className="cell grey">No of UL Bytes: -
                                  </div>
                                </div>
                                <div className="col-1-3 push-1-3">
                                </div>
                                <div className="col-1-3">
                                  <div className="cell grey">No of DL Datagrams: - 
                                  </div>
                                </div>
                                <div className="col-1-3">
                                  <div className="cell grey">No of DL Bytes: 
                                  </div>
                                </div>
                              </div>
                              <p />
                            </div>
                          </div>
                        </div>
                        <p />
                        <h4 style={{marginLeft: 20}}>Advanced Mode</h4>
                        <p>
                        </p><div className="col-lg-12">
                          <div className="panel panel-info" style={{height: 250}}>
                            <div className="panel-body"><p>
                              </p><div className="col-lg-8">
                                <div className="col-1-3">
                                  <div className="cell red">Status: Disconnected
                                  </div>
                                </div>
                              </div>
                              <div className="col-lg-8">
                                <div className="col-1-3">
                                  <div className="cell grey">No of UL Datagrams: -
                                  </div>
                                </div>
                                <div className="col-1-3">
                                  <div className="cell grey">No of UL Bytes: -
                                  </div>
                                </div>
                                <div className="col-1-3 push-1-3">
                                </div>
                                <div className="col-1-3">
                                  <div className="cell grey">No of DL Datagrams: - 
                                  </div>
                                </div>
                                <div className="col-1-3">
                                  <div className="cell grey">No of DL Bytes: 
                                  </div>
                                </div>
                              </div>
                              <p />
                            </div>
                          </div>
                        </div>
                        <p />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              {/* /.panel-body */}
            </div>
            {/* /.panel */}
          </div>
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
module.exports = Mode;