var AppConstants = require('../constants/app-constants');
var AppDispatcher = require('../dispatchers/app-dispatcher');

var AppActions = {
  addItem: function(item){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.ADD_ITEM,
      item: item
    })
  },
  removeItem: function(index){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.REMOVE_ITEM,
      index: index
    })
  },
  increaseItem: function(index){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.INCREASE_ITEM,
      index: index
    })
  },
  decreaseItem: function(index){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.DECREASE_ITEM,
      index: index
    })
  },
  setCommissioning: function(uuid){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.SET_COMMISSIONING,
      uuid: uuid
    })
  },
  setCommissioning: function(uuid){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.SET_TRAFFIC_TEST,
      uuid: uuid
    })
  },
  setCommissioning: function(uuid){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.SET_STANDARD_TRX,
      uuid: uuid
    })
  },
  setCommissioning: function(uuid){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.SET_HEARTBEAT,
      uuid: uuid
    })
  },
  setCommissioning: function(uuid){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.SET_REPORTING_INTERVAL,
      uuid: uuid
    })
  },
  setCommissioning: function(uuid){
    AppDispatcher.handleViewAction({
      actionType: AppConstants.REBOOT,
      uuid: uuid
    })
  }
}

module.exports = AppActions;
