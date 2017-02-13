var app = new Vue({
    el: '#app',
    data: {
        status: {
            code: null,
            text: ''
        },
        headers: [],
        cookies: [],
        responseBody: '',
        method: 'GET',
        mimeType: 'application/json'
    },
    methods: {
        addHeader: function () {

            var self = this;
            self.headers.push({name: '', value: ''});
        },
        deleteHeader: function (index) {

            var self = this;
            self.headers.splice(index, 1)
        },
        addCookie: function () {

            var self = this;
            self.cookies.push({name: '', value: ''});
        },
        deleteCookie: function (index) {

            var self = this;
            self.cookies.splice(index, 1);
        },
        create: function () {

            var self = this;
            if (!checkParams(self)) {
                console.error('checkParams fail');
            } else {
                $.post('/api/response', JSON.stringify({
                        status: self.status,
                        headers: self.headers,
                        cookies: self.cookies,
                        responseBody: self.responseBody,
                        method: self.method,
                        mimeType: self.mimeType
                    }
                ), function (json) {
                    console.info(json);
                    if (json.success) {
                        // open created mockr detail page
                        window.location = '/r/' + json.data;
                    }
                }).fail(function (err) {
                    console.error(err);
                });
            }
        }
    }
});

function checkParams(self) {

    return !(!self.status.code || !self.status.text || !self.method)
}