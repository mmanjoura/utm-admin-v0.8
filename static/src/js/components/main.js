var React = require('react');
var Router = require('react-router-component');
var Template = require('./main-template.js');
var Display = require('./display/display.react');
var Mode = require('./mode/mode.react');

var Catalog = require('./catalog/app-catalog');
var Cart = require('./cart/app-cart');
var Router = require('react-router-component');
var CatalogDetail = require('./product/app-catalogdetail');
var Template = require('./app-template.js');

var Locations = Router.Locations;
var Location  = Router.Location;

var App = React.createClass({
  render:function(){
    return (
      <Template>
        <Locations>
          <Location path="/" handler={Display} />
          <Location path="/mode" handler={Mode} />

          <Location path="/home" handler={Catalog} />
          <Location path="/cart" handler={Cart} />
          <Location path="/item/:item" handler={CatalogDetail} />

        </Locations>
      </Template>
    );
  }
});

module.exports = App;
