var app = angular.module('bank', ['dateInput', 'bankFilters', 'highcharts-ng']);

app.run(function($rootScope) {
    // Broadcast 'reload' events to controllers that need to hit the backend
    $rootScope.$on('reload', function(event, args) {
        $rootScope.$broadcast('reloadBroadcast', args);
    });
    // Broadcast 'refresh graph' events
    $rootScope.$on('refreshGraph', function(event, args) {
        $rootScope.$broadcast('refreshGraphBroadcast', args);
    });
});

// controller for the transactions table
app.controller('transactionsList', function($scope, $http, $filter) {
    $scope.endDate = new Date();
    $scope.startDate = new Date();
    $scope.startDate.setMonth($scope.startDate.getMonth() - 1);

    $scope.dateFilter = function(transaction) {
        var transDate = new Date(transaction.Date);
        return transDate >= $scope.startDate && transDate <= $scope.endDate;
    };

    var emitRefresh = function() {
        $scope.$emit('refreshGraph', {
            'transactions': $filter('filter')($scope.transactions, $scope.dateFilter),
            'startDate': $scope.startDate,
            'endDate': $scope.endDate
        })
    };

    var load = function() {
        $http.get('/transactions').success(function(data) {
            $scope.transactions=data;
            emitRefresh();
        });
    };
    load();
    $scope.$on('reloadBroadcast', load);

    // Trigger a graph refresh on any start or end date change
    $scope.$watch('startDate', function(newVal, oldVal) {
        emitRefresh();
    });
    $scope.$watch('endDate', function(newVal, oldVal) {
        emitRefresh();
    });
});

// controller for transaction creation
app.controller('transaction', function($scope, $http) {
    $scope.transaction = {};
    var now = new Date();
    $scope.transaction.Date = new Date(now.getUTCFullYear(), now.getUTCMonth(), now.getUTCDate(), 0, 0, 0);

    $scope.save = function(transaction) {
        console.log("POSTing: " + JSON.stringify(transaction));
        $http.post('/transaction', transaction).success(function(response) {
            console.log("POST response: " + JSON.stringify(response));
            $scope.transaction = {};
            $scope.$emit('reload', {'transaction': transaction});
        });
    };
});

// controller for the transactions graph
app.controller('transactionsGraph', function($scope, transReportFactory) {
    $scope.chartConfig = {
        options: {
            chart: {
                type: 'area'
            }
        },
        series: [{
            'name': 'Credits',
            'color': '#00B945',
        },
        {
            'name': 'Debits',
            'color': '#FF2C00',
        }],
        title: {
            text: null
        }
    };

    $scope.$on('refreshGraphBroadcast', function(event, args) {
        if (args.transactions) {
            var credits = args.transactions.filter(function(trans) {
                return trans.Amount > 0;
            });
            $scope.chartConfig.series[0].data = transReportFactory.getReport(args.startDate, 
                args.endDate, 
                credits);
            console.log(credits);
            console.log($scope.chartConfig.series[0].data);

            var debits = args.transactions.filter(function(trans) {
                return trans.Amount < 0;
            });
            $scope.chartConfig.series[1].data = transReportFactory.getReport(args.startDate, 
                args.endDate, 
                debits);
            console.log(debits);
            console.log($scope.chartConfig.series[1].data);
        }
    });

});
