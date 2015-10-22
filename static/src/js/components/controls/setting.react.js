var React = require('react');
var Apply = require('./apply.react');
var Reboot = require('./reboot.react');
var Reporting = require('./reporting.react');
var HeartBeat = require('./heartbeat.react');

var Setting = React.createClass({
  render: function() {
    return (
            <div className="navbar-default sidebar" role="navigation">
               <div className="sidebar-nav navbar-collapse">
                 <ul className="nav" id="side-menu">
                  <div className="panel panel-info" style={{height: 200, width: 230, margin: '0 auto'}}>
                    <div className="panel-body"><p>
                    <Reboot />
                        // Reboot
                        <Reporting />
                        //Reporting
                        <HeartBeat />
                        //HeartBeat
                        <Apply />
                       // Apply to all
                      </p>
                    </div>
                  </div>
                 </ul>
                </div>
             </div>
    );
  }
});

module.exports = Setting;