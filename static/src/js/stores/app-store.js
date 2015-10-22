var AppDispatcher = require('../dispatchers/app-dispatcher');
var AppConstants = require('../constants/app-constants');
var assign = require('react/lib/Object.assign');
var EventEmitter = require('events').EventEmitter;

var CHANGE_EVENT = 'change';

pollState(function(state) {

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





/////////////////////////////////////////////////////////////////////////////////

var _catalog = [];

for(var i=1; i<9; i++){
  _catalog.push({
    'id': 'Widget' + i,
    'title':'Widget #' + i,
    'summary': 'This is an awesome widget!',
    'description': 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ducimus, commodi.',
    'cost': i,
    'img': '/assets/product.png'
  });
}

var _cartItems = [];

function _removeItem(index){
  _cartItems[index].inCart = false;
  _cartItems.splice(index, 1);
}

function _increaseItem(index){
  _cartItems[index].qty++;
}

function _decreaseItem(index){
  if(_cartItems[index].qty>1){
    _cartItems[index].qty--;
  }
  else {
    _removeItem(index);
  }
}

function _addItem(item){
  if(!item.inCart){
    item['qty'] = 1;
    item['inCart'] = true;
    _cartItems.push(item);
  }
  else {
    _cartItems.forEach(function(cartItem, i){
      if(cartItem.id===item.id){
        _increaseItem(i);
      }
    });
  }
}

function _cartTotals(){
  var qty =0, total = 0;
  _cartItems.forEach(function(cartItem){
    qty+=cartItem.qty;
    total+=cartItem.qty*cartItem.cost;
  });
  return {'qty': qty, 'total': total};
}

function _setCommissioning(uuid){
}
function _setTrafficTest(uuid){
}
function _setStandardTrx(uuid){
}
function _setHeartBeat(uuid){
}
function _setReportingInterval(uuid){
}
function _reboot(uuid){
}



var AppStore = assign(EventEmitter.prototype, {
  emitChange: function(){
    this.emit(CHANGE_EVENT)
  },

  addChangeListener: function(callback){
    this.on(CHANGE_EVENT, callback)
  },

  removeChangeListener: function(callback){
    this.removeListener(CHANGE_EVENT, callback)
  },

  getCart: function(){
    return _cartItems
  },

  getCatalog: function(){
    return _catalog
  },

  getCartTotals: function(){
    return _cartTotals()
  },

  getUtmsData: function(){
    return state
  },


  dispatcherIndex: AppDispatcher.register(function(payload){
    var action = payload.action; // this is our action from handleViewAction
    switch(action.actionType){
      case AppConstants.ADD_ITEM:
        _addItem(payload.action.item);
        break;

      case AppConstants.REMOVE_ITEM:
        _removeItem(payload.action.index);
        break;

      case AppConstants.INCREASE_ITEM:
        _increaseItem(payload.action.index);
        break;

      case AppConstants.DECREASE_ITEM:
        _decreaseItem(payload.action.index);
        break;

      case AppConstants.SET_COMMISSIONING:
        _setCommissioning(payload.action.index);
        break;
      case AppConstants.SET_TRAFFIC_TEST:
        _setTrafficTest(payload.action.index);
        break;
      case AppConstants.SET_STANDARD_TRX:
        _setStandardTrx(payload.action.index);
        break;
      case AppConstants.SET_HEARTBEAT:
        _setHeartBeat(payload.action.index);
        break;
      case AppConstants.SET_REPORTING_INTERVAL:
        _setReportingInterval(payload.action.index);
        break;
      case AppConstants.REBOOT:
        reboot(payload.action.index);
        break;
    }

    AppStore.emitChange();

    return true;
  })

})

module.exports = AppStore;

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
  "LatestDisplayRow": {}
};

