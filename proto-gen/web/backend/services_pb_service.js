// package: 
// file: backend/services.proto

var backend_services_pb = require("../backend/services_pb");
var backend_messages_pb = require("../backend/messages_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var BackendService = (function () {
  function BackendService() {}
  BackendService.serviceName = "BackendService";
  return BackendService;
}());

BackendService.GetRanking = {
  methodName: "GetRanking",
  service: BackendService,
  requestStream: false,
  responseStream: false,
  requestType: backend_messages_pb.GetRankingRequest,
  responseType: backend_messages_pb.GetRankingResponse
};

BackendService.PostSubmit = {
  methodName: "PostSubmit",
  service: BackendService,
  requestStream: false,
  responseStream: false,
  requestType: backend_messages_pb.PostSubmitRequest,
  responseType: backend_messages_pb.PostSubmitResponse
};

BackendService.GetSubmit = {
  methodName: "GetSubmit",
  service: BackendService,
  requestStream: false,
  responseStream: true,
  requestType: backend_messages_pb.GetSubmitRequest,
  responseType: backend_messages_pb.GetSubmitResponse
};

BackendService.PostLogin = {
  methodName: "PostLogin",
  service: BackendService,
  requestStream: false,
  responseStream: false,
  requestType: backend_messages_pb.PostLoginRequest,
  responseType: backend_messages_pb.PostLoginResponse
};

exports.BackendService = BackendService;

function BackendServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BackendServiceClient.prototype.getRanking = function getRanking(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BackendService.GetRanking, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BackendServiceClient.prototype.postSubmit = function postSubmit(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BackendService.PostSubmit, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BackendServiceClient.prototype.getSubmit = function getSubmit(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(BackendService.GetSubmit, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

BackendServiceClient.prototype.postLogin = function postLogin(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BackendService.PostLogin, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.BackendServiceClient = BackendServiceClient;

var HealthcheckService = (function () {
  function HealthcheckService() {}
  HealthcheckService.serviceName = "HealthcheckService";
  return HealthcheckService;
}());

HealthcheckService.PingUnary = {
  methodName: "PingUnary",
  service: HealthcheckService,
  requestStream: false,
  responseStream: false,
  requestType: backend_messages_pb.PingUnaryRequest,
  responseType: backend_messages_pb.PingUnaryResponse
};

HealthcheckService.PingServerSideStreaming = {
  methodName: "PingServerSideStreaming",
  service: HealthcheckService,
  requestStream: false,
  responseStream: true,
  requestType: backend_messages_pb.PingServerSideStreamingRequest,
  responseType: backend_messages_pb.PingServerSideStreamingResponse
};

exports.HealthcheckService = HealthcheckService;

function HealthcheckServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

HealthcheckServiceClient.prototype.pingUnary = function pingUnary(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(HealthcheckService.PingUnary, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

HealthcheckServiceClient.prototype.pingServerSideStreaming = function pingServerSideStreaming(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(HealthcheckService.PingServerSideStreaming, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.HealthcheckServiceClient = HealthcheckServiceClient;

