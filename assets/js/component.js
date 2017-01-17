Vue.component('mockr-header', {
    template: '\
<div class="container">\
    <nav class="nav">\
    <div class="nav-left">\
    <a class="nav-item" href="/">\
    <h1 class="title">Mockr</h1>\
    </a>\
    </div>\
    <div class="nav-right nav-menu">\
    <a class="nav-item" href="/docs">\
    DOCS\
    </a>\
    <a class="nav-item" href="/r">\
    CREATE\
    </a>\
    <a class="nav-item" href="https://">\
    GITHUB\
    </a>\
    </div>\
    </nav>\
</div>\
    '
});


Vue.component('mockr-footer', {
    template: '\
<footer class="footer">\
    <div class="container">\
    <div class="content has-text-centered">\
    <p>\
    <strong>Mockr</strong> by <a href="https://github.com/xuqingfeng">xuqingfeng</a>.\
    The source code is licensed <a href="http://opensource.org/licenses/mit-license.php">MIT</a>.\
</p>\
</div>\
</div>\
</footer>\
    '
});