var app = angular.module('bank', ['dateInput', 'highcharts-ng']);

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

app.controller('transactionsGraph', function($scope) {
    $scope.chartConfig = {
        options: {
            chart: {
                type: 'area'
            }
        },
        series: [{
            'name': 'Credits',
            'color': '#00B945',
            data: [10, 15, 12, 8, 7]
        },
        {
            'name': 'Debits',
            'color': '#FF2C00',
            data: [1, 2, 3, 4, 5]
        }],
        title: {
            text: null
        }
    };
});
