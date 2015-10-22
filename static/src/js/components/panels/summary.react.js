var React = require('react');

var Summary = React.createClass({
  render:function(){
    return (
     	<div className="col-lg-4">
            <div className="panel panel-info" style={{height: 110, width: 350}}>
              <div className="panel-body">
                <p style={{fontStyle: 'italic'}}>
                  <b>Total Msg:</b> <span className="resetColor">  {/*this.state["LatestDisplayRow"]["TotalMsgs"]*/}</span><br />
                  <b>Total Bytes:</b> {/*this.state["LatestDataVolume"]["UplinkBytes"]*/}<br />
                  <b>Last Msg:</b>    {/*this.state["LatestDisplayRow"]["UlastMsgReceived"]*/}<br />
                </p> 
              </div>
            </div>
          </div>
    );
  }
});

module.exports = Summary;