angular.module('bankFilters', [])
    .filter('utcDate', function($filter) {
        return function(input, format) {
            var date = new Date(input);
            date.setMinutes(date.getMinutes() + date.getTimezoneOffset());
            // 'date' renders in local time by default. The TZ offset is added
            // to convert 'local' to 'UTC' before rendering
            return $filter('date')(date, format);
        }
    });
