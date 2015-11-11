/**
 * Copyright (c) 2014, U-blox.
 * All rights reserved.
 */
var React = require('react');
var Link = require('react-router-component').Link;

var DisplayRow = React.createClass({

    render: function() {

        var rows = [];

        this.props.data.forEach(function(uuid, i) {
        if(uuid["Connection"]["Status"] !== undefined){   

        rows.push(       
    
              <tr className="even gradeC" key={i}>
                        <td style={{width: 15}}>
                          <a tabIndex={-1} href="#/standardtwo"> <b className="fa fa-cogs" /></a><br />
                          <input type="checkbox" style={{width: 15}} /><br />
                          <img src="static/dist/assets/images/green.png" alt="logo" style={{maxWidth: 12}} />
                        </td>
                        <td>
                          <ul className="SmallPadding">
                            <li><b>Uuid:</b> {uuid["LatestDisplayRow"]["Uuid"]}</li>
                            <li><b>Mode:</b> {uuid["LatestDisplayRow"]["Mode"]}</li>
                            <li><b>Name:</b> {uuid["LatestDisplayRow"]["UnitName"]}</li>
                             <li><b>Reporting Interval:</b> {uuid["LatestDisplayRow"]["ReportingInterval"]}</li>
                              <li><b>Heart Beat:</b> {uuid["LatestDisplayRow"]["HeartbeatSeconds"]}</li>
                          </ul>   
                        </td>
                        <td>
                          <ul className="SmallPadding">
                            <li><b>Total Msg:</b> {uuid["LatestDisplayRow"]["UTotalMsgs"]}</li>
                            <li><b>Total Bytes:</b> {uuid["LatestDisplayRow"]["UTotalBytes"]}</li>
                            <li><b>Last Msg RX:</b> {uuid["LatestDisplayRow"]["UlastMsgReceived"]}</li>
                          </ul> 
                        </td>
                        <td className="center">
                          <ul className="SmallPadding">
                            <li><b>Total Msg:</b> {uuid["LatestDisplayRow"]["DTotalMsgs"]}</li>
                            <li><b>Total Bytes:</b> {uuid["LatestDisplayRow"]["DTotalBytes"]}</li>
                            <li><b>Last Msg RX:</b> {uuid["LatestDisplayRow"]["DlastMsgReceived"]}</li>
                          </ul>
                        </td>
                        <td className="center">{uuid["LatestDisplayRow"]["RSRP"]}</td>
                        <td className="center" style={{width: 105}}>
                          <i className="fa fa-floppy-o" /> {uuid["LatestDisplayRow"]["DiskSpaceLeft"]}<br />
                          <i className="fa fa-battery-full" /> {uuid["LatestDisplayRow"]["BatteryLevel"]}
                        </td> 
                      </tr>

                      );

                }
         
   
       });
                return (
                        <div className="row" >
                          <div className="panel panel-default" >
                            <div className="_panel-heading" style={{width:'100%'}}>
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
                                            {rows}
                                    </tbody>
                                  </table>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                );



    }
});

module.exports = DisplayRow;