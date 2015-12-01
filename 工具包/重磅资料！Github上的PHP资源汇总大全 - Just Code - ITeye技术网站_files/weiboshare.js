function open_window(url) {
  function o(){
    if(!window.open(url,'iteye',['toolbar=0,status=0,resizable=1,width=440,height=430,left=',(screen.width-440)/2,',top=',(screen.height-430)/2].join('')))
      window.location.href = url;
  }

  if(/Firefox/.test(navigator.userAgent)) {
    setTimeout(o,0);
  } else {
    o();
  }
}

var WeiboShare = Class.create({
  initialize: function(opts) {
    var imgs = opts['img_scope'].select('img:not(.star):not(.spinner)');
    var img_url = '';
    if(imgs.size() > 0) img_url = imgs[0]['src'];

    this.params = {
      title: opts['title'] || document.title,
      url: encodeURIComponent(opts['url'] || document.location.href),
      pic: encodeURIComponent(img_url)
    }
    this.share_buttons = opts['share_buttons'];

    this.share_buttons.select('a').each(function(link){
      link.observe('click', (function(event){
        var url = WeiboShare.make_share_url(link.readAttribute('data-type'), this.params);
        open_window(url);
        event.stop();
      }).bindAsEventListener(this));
    }, this);
  }
});

WeiboShare.register = function(sites) {
  this.sites = (this.sites || new Hash()).merge(sites);
};

WeiboShare.make_share_url = function(key, params) {
  var site_info = this.sites.get(key);
  params['appkey'] = site_info['appkey'];

  var param_array = [];
  for(var p in params) {
    param_array.push(p + '=' + params[p]);
  }

  return (site_info['url'] + '?' + param_array.join('&'));
}

WeiboShare.register({
  sina: {
    url: 'http://service.weibo.com/share/share.php',
    appkey: '3842512498'
  },
  qq: {
    url: 'http://v.t.qq.com/share/share.php',
    appkey: '050a47d9d5d848029e1de3198d2abcda'
  }
});
