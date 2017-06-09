function formatTime(date) {
  var myDate = new Date(date*1000);
  var year = myDate.getFullYear()
  var month = myDate.getMonth() + 1
  var day = myDate.getDate()

  var hour = myDate.getHours()
  var minute = myDate.getMinutes()
  var second = myDate.getSeconds()


  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}
function formatNumber(n) {
  n = n.toString()
  return n[1] ? n : '0' + n
}
function refresh_iframe(){
    window.location.reload();
}