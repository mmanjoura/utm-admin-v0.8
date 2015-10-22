var React = require('react');
var Link = require('react-router-component').Link;

var Mode = React.createClass({
  render:function(){
    return (
    	<div className="col-lg-4">
            <div className="panel" style={{height: 60, width: 180, marginTop: 100}}>
              <div className="panel-body"><p>
                  <a tabIndex={-1} href="#/standard"> <b className="fa fa-cogs  fa-2x" /></a> 
                  <b style={{float: 'right'}}>Configure Ticked</b><br />
                <Link href="/mode">Change Unit Mode</Link>
                </p>
              </div>
            </div>
          </div>
     
    );
  }
});

module.exports = Mode;
