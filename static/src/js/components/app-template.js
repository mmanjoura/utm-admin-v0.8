var React = require('react');
var Setting = require('./controls/setting.react');


var Template = React.createClass({

	render:function(){
	return (
	  <div className="container">
	  {/*<Setting />*/}
	
	    {this.props.children}
	  </div>
	);
	}
});

module.exports = Template;
