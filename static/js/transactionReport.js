var app = angular.module('bank');

// A service for generating graphable accumulation reports from lists of credits & debits
app.factory('transReportFactory', function() {
    return {
        // Return an array containing one element for each day between startDate
        // and endDate (inclusive). Each element is the sum of transactions that
        // occurred previously
        'getReport': function(startDate, endDate, transactions) {
            console.log(startDate);
            console.log(endDate);
            var report = [];
            var sum = 0;
            var dateIterator = new Date(startDate.getTime());
            var transactionsIter = transactions.slice(0);

            //TODO: off-by-one date error (extra array entry at the head)
            while (dateIterator <= endDate) {
                while (transactionsIter.length > 0 && this.sameDay(dateIterator, 
                    new Date(transactionsIter[0].Date))) {
                    sum += transactionsIter.shift().Amount;
                }

                report.push(sum);
                dateIterator.setDate(dateIterator.getDate() + 1);
            }

            return report;
        },
        'sameDay': function(dateA, dateB) {
            return dateA.getDate() == dateB.getDate()
                && dateA.getMonth() == dateB.getMonth()
                && dateA.getYear() == dateB.getYear();
        }
    }
});
