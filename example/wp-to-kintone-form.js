(function ($) {
  $(function () {
    if ('undefined' !== typeof kintone_error) {
      for (var e in kintone_error.errors) {
        var id = e.replace(/\]?.values?$/, '').replace(/^record(\.|\[)/, '');
        var message = kintone_error.errors[e].messages[0];
        if ($('#' + id)[0].nodeName == 'DIV') {
          $('label', '#' + id).addClass('kintone-alert-label');
          $('#' + id).append($('<div class="kintone-error-message" />').text(message));
        } else {
          $('#' + id).addClass('kintone-alert-input');
          $('#' + id).parent().append($('<div class="kintone-error-message" />').text(message));
        }
      }
    }

    $(".form-control.date").datetimepicker({
      lang: 'ja',
      i18n: {
        ja: {
          months: ["1\u6708", "2\u6708", "3\u6708", "4\u6708", "5\u6708", "6\u6708", "7\u6708", "8\u6708", "9\u6708", "10\u6708", "11\u6708", "12\u6708"],
          dayOfWeek: ["\u65e5", "\u6708", "\u706b", "\u6c34", "\u6728", "\u91d1", "\u571f"]
        }
      },
      scrollInput: false,
      timepicker: false,
      format: 'Y-m-d'
    });

    $(".form-control.time").datetimepicker({
      lang: 'ja',
      i18n: {
        ja: {
          months: ["1\u6708", "2\u6708", "3\u6708", "4\u6708", "5\u6708", "6\u6708", "7\u6708", "8\u6708", "9\u6708", "10\u6708", "11\u6708", "12\u6708"],
          dayOfWeek: ["\u65e5", "\u6708", "\u706b", "\u6c34", "\u6728", "\u91d1", "\u571f"]
        }
      },
      scrollInput: false,
      datepicker: false,
      format: 'H:i'
    });

    $(".form-control.datetime").datetimepicker({
      lang: 'ja',
      i18n: {
        ja: {
          months: ["1\u6708", "2\u6708", "3\u6708", "4\u6708", "5\u6708", "6\u6708", "7\u6708", "8\u6708", "9\u6708", "10\u6708", "11\u6708", "12\u6708"],
          dayOfWeek: ["\u65e5", "\u6708", "\u706b", "\u6c34", "\u6728", "\u91d1", "\u571f"]
        }
      },
      scrollInput: false,
      format: 'c'
    });

    Date.parseDate = function (input, format) {
      return Date.parse(input);
    };

    Date.prototype.dateFormat = function (format) {
      var offsetmin = this.getTimezoneOffset();
      var offset = 0 - (offsetmin / 60);
      if (Math.abs(offset) < 0) {
        offset = (10 > Math.abs(offset)) ? '-0' + Math.abs(offset) + ':00' : '-' + Math.abs(offset) + ':00';
      } else {
        offset = (10 > Math.abs(offset)) ? '+0' + Math.abs(offset) + ':00' : '+' + Math.abs(offset) + ':00';
      }

      var y = this.getFullYear();
      var m = (10 > (this.getMonth() + 1)) ? '0' + (this.getMonth() + 1) : this.getMonth() + 1;
      var d = (10 > this.getDate()) ? '0' + this.getDate() : this.getDate();
      var h = (10 > this.getHours()) ? '0' + this.getHours() : this.getHours();
      var i = (10 > this.getMinutes()) ? '0' + this.getMinutes() : this.getMinutes();

      switch (format) {
        case 'Y-m-d':
          return y + '-' + m + '-' + d;
        case 'H:i':
          return h + ':' + i;
        case 'c':
          return y + '-' + m + '-' + d + 'T' + h + ':' + i + offset;
      }
      // or default format
      return this.getDate() + '.' + (this.getMonth() + 1) + '.' + this.getFullYear();
    };
  });
})(jQuery);
