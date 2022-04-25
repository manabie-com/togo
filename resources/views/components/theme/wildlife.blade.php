<style type="text/css">
/*theme-ufala file*/
html body{background-color: #F9F9F9; color: #243648;}
/*nav-bar*/
html body.fixed-navbar { padding-top: 0rem; }
html body .content-wrapper {padding: 48px}
nav.header-navbar, #menu-toggle-ext {display: none!important;}
/*set fonts*/
html body, .main-menu .navigation,
.main-menu .navigation .navigation-header,
select, button, textarea, text, input {font-family: 'Roboto', sans-serif!important;}

/*pagination*/
.pagination li.page-first a i::before{content: url("/images/menu/page-first.png");}
.pagination li.page-last a i::before{content: url("/images/menu/page-last.png");}
.pagination li.page-next a i::before{content: url("/images/menu/page-next.png");}
.pagination li.page-previous a i::before{content: url("/images/menu/page-previous.png");}

/*--VENDER--*/
/*Kint*/
.kint-rich {bottom: 100px;}
.kint-rich.kint-folder{top: 0px!important;}

/*Xdebug*/
.xdebug-error.xe-xdebug {z-index: 4000;font-family: monospace;top: 0;left: 0;font-size: 16px;}
</style>
@push('scripts')
<script type="text/javascript">
    /*MENU*/
    $('.main-menu.menu-light .navigation > li.open:has("ul") > a').css('color','inherit');
    $('.main-menu.menu-light .navigation > li.open:has("ul") > a .icon >img').css('filter','none');
</script>
@endpush
