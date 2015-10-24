/**
 * Copyright (c) 2014, U-blox.
 * All rights reserved.
 */
var React = require('react');
var Link = require('react-router-component').Link;

var DisplayRow = React.createClass({

    render: function() {

        return (
            <table className="table table-striped table-bordered table-hover" id="dataTables-example">
                <thead>
                    <tr className="info">
                        <th> <input type="checkbox" /> All</th>
                        <th>Name - UUID</th>
                        <th>Upstream</th>
                        <th>Downsteam</th>
                        <th>RSRP </th>
                        <th>
                      Battery - <i className="fa fa-floppy-o"></i>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr className="even gradeA" >
                        <td>
                            <div>
                                                                      
                                    <a href="#/standard"  className="dropdown-toggle js-activated" >
                                        <b className="fa fa-cogs"></b>
                                    </a>
                            </div>
                            
                            <div>
                              
                                        <input type="checkbox" />
                            </div>
                            <div>
                    
                                        <img src={this.props.item} width="10px" height="10px" alt="logo" />
                         
                            </div>

                        </td>
                        <td>
                                <ul>
                                    <li><b>Uuid:</b> {this.props.item}</li>
                                    <li><b>Mode:</b> {this.props.item}</li>
                                    <li><b>Name:</b> {this.props.item}</li>
                                </ul>   
                        </td>
                        <td>
                               <ul>
                                    <li><b>Total Msg:</b> {this.props.item}</li>
                                    <li><b>Total Bytes:</b> {this.props.item}</li>
                                    <li><b>Last Msg RX:</b> {this.props.item}</li>
                                </ul> 
                        </td>
                        <td className="center">
                                <ul>
                                    <li><b>Total Msg:</b> {this.props.item}</li>
                                    <li><b>Total Bytes:</b> {this.props.item}</li>
                                    <li><b>Last Msg RX:</b>{this.props.item}</li>
                                </ul>
                        </td>
                        <td className="center">{this.props.item}</td>
                        <td className="center">{this.props.item ? "" : "N/A"} - 
                                               {this.props.item ? "" : "N/A"}</td> 
                    </tr>
                </tbody>
            </table>
        );
    }
});

module.exports = DisplayRow;