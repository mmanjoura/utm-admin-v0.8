var React = require('react');


var Configure = React.createClass({
  render:function(){
    return (
    	<div className="col-lg-4">
            <div className="panel" style={{height: 60, width: 200, marginTop: 100}}>
              <div className="panel-body">
                <p>
                  <a tabIndex={-1} href="#/standard"> <b className="fa fa-cogs  fa-2x" /></a> 
                  <b style={{float: 'right'}}>Configures Ticked</b><br />
                </p>
              </div>
            </div>
          </div>
     
    );
  }
});

module.exports = Configure;
