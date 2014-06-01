var app = angular.module('bank', ['dateInput']);

app.run(function($rootScope) {
    // Broadcast 'reload' events to controllers that need to hit the backend
    $rootScope.$on('reload', function(event, args) {
        $rootScope.$broadcast('reloadBroadcast', args);
    });
});

app.controller('transactionsList', function($scope, $http) {
    var load = function() {
        $http.get('/transactions').success(function(data) {
            $scope.transactions=data;
        });
    };
    load();
    $scope.$on('reloadBroadcast', load);

    $scope.endDate = new Date();
    $scope.startDate = new Date();
    $scope.startDate.setMonth($scope.startDate.getMonth() - 1);

    $scope.dateFilter = function(transaction) {
        var transDate = new Date(transaction.Date);
        return transDate >= $scope.startDate && transDate <= $scope.endDate;
    };
});

app.controller('transaction', function($scope, $http) {
    $scope.transaction = {};
    $scope.transaction.Date = new Date();

    $scope.save = function(transaction) {
        console.log("POSTing: " + JSON.stringify(transaction));
        $http.post('/transaction', transaction).success(function(response) {
            console.log("POST response: " + JSON.stringify(response));
            $scope.transaction = {};
            $scope.$emit('reload', {'transaction': transaction});
        });
    };
});
