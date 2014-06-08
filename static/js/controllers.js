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

    var refresh = function() {
        var filtered = $filter('filter')($scope.transactions, $scope.dateFilter);
        $scope.$emit('refreshGraph', {
            'transactions': filtered,
            'startDate': $scope.startDate,
            'endDate': $scope.endDate
        })

        var total = 0;
        angular.forEach(filtered, function(trans) {
            total += trans.Amount;
        });
        $scope.total = total;
    };

    var load = function() {
        $http.get('/transactions').success(function(data) {
            $scope.transactions=data;
            refresh();
        });
    };
    load();
    $scope.$on('reloadBroadcast', load);

    // Trigger a graph refresh on any start or end date change
    $scope.$watch('startDate', function(newVal, oldVal) {
        refresh();
    });
    $scope.$watch('endDate', function(newVal, oldVal) {
        refresh();
    });

    $scope.delete = function(transaction) {
        console.log('DELETE-ing: ' + JSON.stringify(transaction));
        $http.delete('/transaction/' + transaction.Id).success(function(response) {
            console.log('DELETE response: ' + JSON.stringify(response));
            $scope.$emit('reload');
        });
    };
});

// controller for transaction creation
app.controller('transaction', function($scope, $http, $filter) {
    var resetTransaction = function() {
        $scope.transaction = {};
        var now = new Date();
        $scope.transaction.Date = new Date(now.getUTCFullYear(), now.getUTCMonth(), now.getUTCDate(), 0, 0, 0);
    };
    resetTransaction();

    $scope.save = function(transaction) {
        var newTrans = {
            'Date': transaction.Date,
            'Amount': $filter('fromDecimal')(transaction.Amount),
            'Note': transaction.Note
        };

        console.log("POSTing: " + JSON.stringify(newTrans));
        $http.post('/transaction', newTrans).success(function(response) {
            console.log("POST response: " + JSON.stringify(response));
            resetTransaction();
            $scope.$emit('reload', {'transaction': newTrans});
        });
    };
});

// controller for the transactions graph
app.controller('transactionsGraph', function($scope, $filter, transReportFactory) {
    $scope.chartConfig = {
        options: {
            chart: {
                type: 'area'
            },
            tooltip: {
              pointFormat: '<span style="color:{series.color}">\u25CF</span> {series.name}: <b>${point.y:.2f}</b><br/>'
            }
        },
        series: [{
            'name': 'Credits',
            'color': '#00B945',
            'pointInterval': 24 * 3600 * 1000
        },
        {
            'name': 'Debits',
            'color': '#FF2C00',
            'pointInterval': 24 * 3600 * 1000
        }],
        title: {
            text: null
        },
        xAxis: {
            type: 'datetime'
        },
        yAxis: {
            title: {
                text: 'Dollars'
            }
        }
    };

    $scope.$on('refreshGraphBroadcast', function(event, args) {
        // Don't attempt to render the chart for >= 5 years of data
        var dateDiff = Math.abs(args.startDate.getUTCFullYear() - args.endDate.getUTCFullYear());
        if (dateDiff >= 5) {
            $scope.chartConfig.title.text = 'Date range too wide to render';
            $scope.chartConfig.series[0].data = null;
            $scope.chartConfig.series[1].data = null;
        } else if (args.transactions) {
            $scope.chartConfig.title.text = null;
            var credits = $filter('filter')(args.transactions, function(trans) {
                return trans.Amount > 0;
            });
            var creditsReport = transReportFactory.getReport(args.startDate, 
                args.endDate, 
                credits);
            $scope.chartConfig.series[0].data = creditsReport.map($filter('toDecimal'));

            var debits = $filter('filter')(args.transactions, function(trans) {
                return trans.Amount < 0;
            });
            var debitsReport = transReportFactory.getReport(args.startDate, 
                args.endDate, 
                debits);
            $scope.chartConfig.series[1].data = debitsReport
                .map($filter('toDecimal'))
                .map(function(val) {
                    return Math.abs(val);
                });
        }

        $scope.chartConfig.series[0].pointStart = args.startDate.getTime();
        $scope.chartConfig.series[1].pointStart = $scope.chartConfig.series[0].pointStart;
    });

});
