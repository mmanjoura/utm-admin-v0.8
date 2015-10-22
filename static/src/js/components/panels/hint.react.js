var React = require('react');

var Hint = React.createClass({
  render:function(){
    return (
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
     
    );
  }
});

module.exports = Hint;