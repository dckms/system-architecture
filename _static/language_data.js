/*
 * language_data.js
 * ~~~~~~~~~~~~~~~~
 *
 * This script contains the language-specific data used by searchtools.js,
 * namely the list of stopwords, stemmer, scorer and splitter.
 *
 * :copyright: Copyright 2007-2021 by the Sphinx team, see AUTHORS.
 * :license: BSD, see LICENSE for details.
 *
 */

var stopwords = ["\u0430","\u0431\u0435\u0437","\u0431\u043e\u043b\u0435\u0435","\u0431\u043e\u043b\u044c\u0448\u0435","\u0431\u0443\u0434\u0435\u0442","\u0431\u0443\u0434\u0442\u043e","\u0431\u044b","\u0431\u044b\u043b","\u0431\u044b\u043b\u0430","\u0431\u044b\u043b\u0438","\u0431\u044b\u043b\u043e","\u0431\u044b\u0442\u044c","\u0432","\u0432\u0430\u043c","\u0432\u0430\u0441","\u0432\u0434\u0440\u0443\u0433","\u0432\u0435\u0434\u044c","\u0432\u043e","\u0432\u043e\u0442","\u0432\u043f\u0440\u043e\u0447\u0435\u043c","\u0432\u0441\u0435","\u0432\u0441\u0435\u0433\u0434\u0430","\u0432\u0441\u0435\u0433\u043e","\u0432\u0441\u0435\u0445","\u0432\u0441\u044e","\u0432\u044b","\u0433\u0434\u0435","\u0433\u043e\u0432\u043e\u0440\u0438\u043b","\u0434\u0430","\u0434\u0430\u0436\u0435","\u0434\u0432\u0430","\u0434\u043b\u044f","\u0434\u043e","\u0434\u0440\u0443\u0433\u043e\u0439","\u0435\u0433\u043e","\u0435\u0435","\u0435\u0439","\u0435\u043c\u0443","\u0435\u0441\u043b\u0438","\u0435\u0441\u0442\u044c","\u0435\u0449\u0435","\u0436","\u0436\u0435","\u0436\u0438\u0437\u043d\u044c","\u0437\u0430","\u0437\u0430\u0447\u0435\u043c","\u0437\u0434\u0435\u0441\u044c","\u0438","\u0438\u0437","\u0438\u043b\u0438","\u0438\u043c","\u0438\u043d\u043e\u0433\u0434\u0430","\u0438\u0445","\u043a","\u043a\u0430\u0436\u0435\u0442\u0441\u044f","\u043a\u0430\u043a","\u043a\u0430\u043a\u0430\u044f","\u043a\u0430\u043a\u043e\u0439","\u043a\u043e\u0433\u0434\u0430","\u043a\u043e\u043d\u0435\u0447\u043d\u043e","\u043a\u0442\u043e","\u043a\u0443\u0434\u0430","\u043b\u0438","\u043b\u0443\u0447\u0448\u0435","\u043c\u0435\u0436\u0434\u0443","\u043c\u0435\u043d\u044f","\u043c\u043d\u0435","\u043c\u043d\u043e\u0433\u043e","\u043c\u043e\u0436\u0435\u0442","\u043c\u043e\u0436\u043d\u043e","\u043c\u043e\u0439","\u043c\u043e\u044f","\u043c\u044b","\u043d\u0430","\u043d\u0430\u0434","\u043d\u0430\u0434\u043e","\u043d\u0430\u043a\u043e\u043d\u0435\u0446","\u043d\u0430\u0441","\u043d\u0435","\u043d\u0435\u0433\u043e","\u043d\u0435\u0435","\u043d\u0435\u0439","\u043d\u0435\u043b\u044c\u0437\u044f","\u043d\u0435\u0442","\u043d\u0438","\u043d\u0438\u0431\u0443\u0434\u044c","\u043d\u0438\u043a\u043e\u0433\u0434\u0430","\u043d\u0438\u043c","\u043d\u0438\u0445","\u043d\u0438\u0447\u0435\u0433\u043e","\u043d\u043e","\u043d\u0443","\u043e","\u043e\u0431","\u043e\u0434\u0438\u043d","\u043e\u043d","\u043e\u043d\u0430","\u043e\u043d\u0438","\u043e\u043f\u044f\u0442\u044c","\u043e\u0442","\u043f\u0435\u0440\u0435\u0434","\u043f\u043e","\u043f\u043e\u0434","\u043f\u043e\u0441\u043b\u0435","\u043f\u043e\u0442\u043e\u043c","\u043f\u043e\u0442\u043e\u043c\u0443","\u043f\u043e\u0447\u0442\u0438","\u043f\u0440\u0438","\u043f\u0440\u043e","\u0440\u0430\u0437","\u0440\u0430\u0437\u0432\u0435","\u0441","\u0441\u0430\u043c","\u0441\u0432\u043e\u044e","\u0441\u0435\u0431\u0435","\u0441\u0435\u0431\u044f","\u0441\u0435\u0433\u043e\u0434\u043d\u044f","\u0441\u0435\u0439\u0447\u0430\u0441","\u0441\u043a\u0430\u0437\u0430\u043b","\u0441\u043a\u0430\u0437\u0430\u043b\u0430","\u0441\u043a\u0430\u0437\u0430\u0442\u044c","\u0441\u043e","\u0441\u043e\u0432\u0441\u0435\u043c","\u0442\u0430\u043a","\u0442\u0430\u043a\u043e\u0439","\u0442\u0430\u043c","\u0442\u0435\u0431\u044f","\u0442\u0435\u043c","\u0442\u0435\u043f\u0435\u0440\u044c","\u0442\u043e","\u0442\u043e\u0433\u0434\u0430","\u0442\u043e\u0433\u043e","\u0442\u043e\u0436\u0435","\u0442\u043e\u043b\u044c\u043a\u043e","\u0442\u043e\u043c","\u0442\u043e\u0442","\u0442\u0440\u0438","\u0442\u0443\u0442","\u0442\u044b","\u0443","\u0443\u0436","\u0443\u0436\u0435","\u0445\u043e\u0440\u043e\u0448\u043e","\u0445\u043e\u0442\u044c","\u0447\u0435\u0433\u043e","\u0447\u0435\u043b\u043e\u0432\u0435\u043a","\u0447\u0435\u043c","\u0447\u0435\u0440\u0435\u0437","\u0447\u0442\u043e","\u0447\u0442\u043e\u0431","\u0447\u0442\u043e\u0431\u044b","\u0447\u0443\u0442\u044c","\u044d\u0442\u0438","\u044d\u0442\u043e\u0433\u043e","\u044d\u0442\u043e\u0439","\u044d\u0442\u043e\u043c","\u044d\u0442\u043e\u0442","\u044d\u0442\u0443","\u044f"];


/* Non-minified version is copied as a separate JS file, is available */
BaseStemmer=function(){this.setCurrent=function(r){this.current=r;this.cursor=0;this.limit=this.current.length;this.limit_backward=0;this.bra=this.cursor;this.ket=this.limit};this.getCurrent=function(){return this.current};this.copy_from=function(r){this.current=r.current;this.cursor=r.cursor;this.limit=r.limit;this.limit_backward=r.limit_backward;this.bra=r.bra;this.ket=r.ket};this.in_grouping=function(r,t,i){if(this.cursor>=this.limit)return false;var s=this.current.charCodeAt(this.cursor);if(s>i||s<t)return false;s-=t;if((r[s>>>3]&1<<(s&7))==0)return false;this.cursor++;return true};this.in_grouping_b=function(r,t,i){if(this.cursor<=this.limit_backward)return false;var s=this.current.charCodeAt(this.cursor-1);if(s>i||s<t)return false;s-=t;if((r[s>>>3]&1<<(s&7))==0)return false;this.cursor--;return true};this.out_grouping=function(r,t,i){if(this.cursor>=this.limit)return false;var s=this.current.charCodeAt(this.cursor);if(s>i||s<t){this.cursor++;return true}s-=t;if((r[s>>>3]&1<<(s&7))==0){this.cursor++;return true}return false};this.out_grouping_b=function(r,t,i){if(this.cursor<=this.limit_backward)return false;var s=this.current.charCodeAt(this.cursor-1);if(s>i||s<t){this.cursor--;return true}s-=t;if((r[s>>>3]&1<<(s&7))==0){this.cursor--;return true}return false};this.eq_s=function(r){if(this.limit-this.cursor<r.length)return false;if(this.current.slice(this.cursor,this.cursor+r.length)!=r){return false}this.cursor+=r.length;return true};this.eq_s_b=function(r){if(this.cursor-this.limit_backward<r.length)return false;if(this.current.slice(this.cursor-r.length,this.cursor)!=r){return false}this.cursor-=r.length;return true};this.find_among=function(r){var t=0;var i=r.length;var s=this.cursor;var e=this.limit;var h=0;var u=0;var n=false;while(true){var c=t+(i-t>>>1);var a=0;var f=h<u?h:u;var l=r[c];var o;for(o=f;o<l[0].length;o++){if(s+f==e){a=-1;break}a=this.current.charCodeAt(s+f)-l[0].charCodeAt(o);if(a!=0)break;f++}if(a<0){i=c;u=f}else{t=c;h=f}if(i-t<=1){if(t>0)break;if(i==t)break;if(n)break;n=true}}do{var l=r[t];if(h>=l[0].length){this.cursor=s+l[0].length;if(l.length<4)return l[2];var v=l[3](this);this.cursor=s+l[0].length;if(v)return l[2]}t=l[1]}while(t>=0);return 0};this.find_among_b=function(r){var t=0;var i=r.length;var s=this.cursor;var e=this.limit_backward;var h=0;var u=0;var n=false;while(true){var c=t+(i-t>>1);var a=0;var f=h<u?h:u;var l=r[c];var o;for(o=l[0].length-1-f;o>=0;o--){if(s-f==e){a=-1;break}a=this.current.charCodeAt(s-1-f)-l[0].charCodeAt(o);if(a!=0)break;f++}if(a<0){i=c;u=f}else{t=c;h=f}if(i-t<=1){if(t>0)break;if(i==t)break;if(n)break;n=true}}do{var l=r[t];if(h>=l[0].length){this.cursor=s-l[0].length;if(l.length<4)return l[2];var v=l[3](this);this.cursor=s-l[0].length;if(v)return l[2]}t=l[1]}while(t>=0);return 0};this.replace_s=function(r,t,i){var s=i.length-(t-r);this.current=this.current.slice(0,r)+i+this.current.slice(t);this.limit+=s;if(this.cursor>=t)this.cursor+=s;else if(this.cursor>r)this.cursor=r;return s};this.slice_check=function(){if(this.bra<0||this.bra>this.ket||this.ket>this.limit||this.limit>this.current.length){return false}return true};this.slice_from=function(r){var t=false;if(this.slice_check()){this.replace_s(this.bra,this.ket,r);t=true}return t};this.slice_del=function(){return this.slice_from("")};this.insert=function(r,t,i){var s=this.replace_s(r,t,i);if(r<=this.bra)this.bra+=s;if(r<=this.ket)this.ket+=s};this.slice_to=function(){var r="";if(this.slice_check()){r=this.current.slice(this.bra,this.ket)}return r};this.assign_to=function(){return this.current.slice(0,this.limit)}};
RussianStemmer=function(){var r=new BaseStemmer;var e=[["в",-1,1],["ив",0,2],["ыв",0,2],["вши",-1,1],["ивши",3,2],["ывши",3,2],["вшись",-1,1],["ившись",6,2],["ывшись",6,2]];var i=[["ее",-1,1],["ие",-1,1],["ое",-1,1],["ые",-1,1],["ими",-1,1],["ыми",-1,1],["ей",-1,1],["ий",-1,1],["ой",-1,1],["ый",-1,1],["ем",-1,1],["им",-1,1],["ом",-1,1],["ым",-1,1],["его",-1,1],["ого",-1,1],["ему",-1,1],["ому",-1,1],["их",-1,1],["ых",-1,1],["ею",-1,1],["ою",-1,1],["ую",-1,1],["юю",-1,1],["ая",-1,1],["яя",-1,1]];var u=[["ем",-1,1],["нн",-1,1],["вш",-1,1],["ивш",2,2],["ывш",2,2],["щ",-1,1],["ющ",5,1],["ующ",6,2]];var s=[["сь",-1,1],["ся",-1,1]];var a=[["ла",-1,1],["ила",0,2],["ыла",0,2],["на",-1,1],["ена",3,2],["ете",-1,1],["ите",-1,2],["йте",-1,1],["ейте",7,2],["уйте",7,2],["ли",-1,1],["или",10,2],["ыли",10,2],["й",-1,1],["ей",13,2],["уй",13,2],["л",-1,1],["ил",16,2],["ыл",16,2],["ем",-1,1],["им",-1,2],["ым",-1,2],["н",-1,1],["ен",22,2],["ло",-1,1],["ило",24,2],["ыло",24,2],["но",-1,1],["ено",27,2],["нно",27,1],["ет",-1,1],["ует",30,2],["ит",-1,2],["ыт",-1,2],["ют",-1,1],["уют",34,2],["ят",-1,2],["ны",-1,1],["ены",37,2],["ть",-1,1],["ить",39,2],["ыть",39,2],["ешь",-1,1],["ишь",-1,2],["ю",-1,2],["ую",44,2]];var t=[["а",-1,1],["ев",-1,1],["ов",-1,1],["е",-1,1],["ие",3,1],["ье",3,1],["и",-1,1],["еи",6,1],["ии",6,1],["ами",6,1],["ями",6,1],["иями",10,1],["й",-1,1],["ей",12,1],["ией",13,1],["ий",12,1],["ой",12,1],["ам",-1,1],["ем",-1,1],["ием",18,1],["ом",-1,1],["ям",-1,1],["иям",21,1],["о",-1,1],["у",-1,1],["ах",-1,1],["ях",-1,1],["иях",26,1],["ы",-1,1],["ь",-1,1],["ю",-1,1],["ию",30,1],["ью",30,1],["я",-1,1],["ия",33,1],["ья",33,1]];var c=[["ост",-1,1],["ость",-1,1]];var f=[["ейше",-1,1],["н",-1,2],["ейш",-1,1],["ь",-1,3]];var l=[33,65,8,232];var o=0;var n=0;function b(){n=r.limit;o=r.limit;var e=r.cursor;r:{e:while(true){i:{if(!r.in_grouping(l,1072,1103)){break i}break e}if(r.cursor>=r.limit){break r}r.cursor++}n=r.cursor;e:while(true){i:{if(!r.out_grouping(l,1072,1103)){break i}break e}if(r.cursor>=r.limit){break r}r.cursor++}e:while(true){i:{if(!r.in_grouping(l,1072,1103)){break i}break e}if(r.cursor>=r.limit){break r}r.cursor++}e:while(true){i:{if(!r.out_grouping(l,1072,1103)){break i}break e}if(r.cursor>=r.limit){break r}r.cursor++}o=r.cursor}r.cursor=e;return true}function _(){if(!(o<=r.cursor)){return false}return true}function k(){var i;r.ket=r.cursor;i=r.find_among_b(e);if(i==0){return false}r.bra=r.cursor;switch(i){case 1:r:{var u=r.limit-r.cursor;e:{if(!r.eq_s_b("а")){break e}break r}r.cursor=r.limit-u;if(!r.eq_s_b("я")){return false}}if(!r.slice_del()){return false}break;case 2:if(!r.slice_del()){return false}break}return true}function m(){r.ket=r.cursor;if(r.find_among_b(i)==0){return false}r.bra=r.cursor;if(!r.slice_del()){return false}return true}function v(){var e;if(!m()){return false}var i=r.limit-r.cursor;r:{r.ket=r.cursor;e=r.find_among_b(u);if(e==0){r.cursor=r.limit-i;break r}r.bra=r.cursor;switch(e){case 1:e:{var s=r.limit-r.cursor;i:{if(!r.eq_s_b("а")){break i}break e}r.cursor=r.limit-s;if(!r.eq_s_b("я")){r.cursor=r.limit-i;break r}}if(!r.slice_del()){return false}break;case 2:if(!r.slice_del()){return false}break}}return true}function d(){r.ket=r.cursor;if(r.find_among_b(s)==0){return false}r.bra=r.cursor;if(!r.slice_del()){return false}return true}function g(){var e;r.ket=r.cursor;e=r.find_among_b(a);if(e==0){return false}r.bra=r.cursor;switch(e){case 1:r:{var i=r.limit-r.cursor;e:{if(!r.eq_s_b("а")){break e}break r}r.cursor=r.limit-i;if(!r.eq_s_b("я")){return false}}if(!r.slice_del()){return false}break;case 2:if(!r.slice_del()){return false}break}return true}function w(){r.ket=r.cursor;if(r.find_among_b(t)==0){return false}r.bra=r.cursor;if(!r.slice_del()){return false}return true}function h(){r.ket=r.cursor;if(r.find_among_b(c)==0){return false}r.bra=r.cursor;if(!_()){return false}if(!r.slice_del()){return false}return true}function q(){var e;r.ket=r.cursor;e=r.find_among_b(f);if(e==0){return false}r.bra=r.cursor;switch(e){case 1:if(!r.slice_del()){return false}r.ket=r.cursor;if(!r.eq_s_b("н")){return false}r.bra=r.cursor;if(!r.eq_s_b("н")){return false}if(!r.slice_del()){return false}break;case 2:if(!r.eq_s_b("н")){return false}if(!r.slice_del()){return false}break;case 3:if(!r.slice_del()){return false}break}return true}this.stem=function(){var e=r.cursor;r:{while(true){var i=r.cursor;e:{i:while(true){var u=r.cursor;u:{r.bra=r.cursor;if(!r.eq_s("ё")){break u}r.ket=r.cursor;r.cursor=u;break i}r.cursor=u;if(r.cursor>=r.limit){break e}r.cursor++}if(!r.slice_from("е")){return false}continue}r.cursor=i;break}}r.cursor=e;b();r.limit_backward=r.cursor;r.cursor=r.limit;if(r.cursor<n){return false}var s=r.limit_backward;r.limit_backward=n;var a=r.limit-r.cursor;r:{e:{var t=r.limit-r.cursor;i:{if(!k()){break i}break e}r.cursor=r.limit-t;var c=r.limit-r.cursor;i:{if(!d()){r.cursor=r.limit-c;break i}}i:{var f=r.limit-r.cursor;u:{if(!v()){break u}break i}r.cursor=r.limit-f;u:{if(!g()){break u}break i}r.cursor=r.limit-f;if(!w()){break r}}}}r.cursor=r.limit-a;var l=r.limit-r.cursor;r:{r.ket=r.cursor;if(!r.eq_s_b("и")){r.cursor=r.limit-l;break r}r.bra=r.cursor;if(!r.slice_del()){return false}}var o=r.limit-r.cursor;h();r.cursor=r.limit-o;var _=r.limit-r.cursor;q();r.cursor=r.limit-_;r.limit_backward=s;r.cursor=r.limit_backward;return true};this["stemWord"]=function(e){r.setCurrent(e);this.stem();return r.getCurrent()}};
Stemmer = RussianStemmer;



var splitChars = (function() {
    var result = {};
    var singles = [96, 180, 187, 191, 215, 247, 749, 885, 903, 907, 909, 930, 1014, 1648,
         1748, 1809, 2416, 2473, 2481, 2526, 2601, 2609, 2612, 2615, 2653, 2702,
         2706, 2729, 2737, 2740, 2857, 2865, 2868, 2910, 2928, 2948, 2961, 2971,
         2973, 3085, 3089, 3113, 3124, 3213, 3217, 3241, 3252, 3295, 3341, 3345,
         3369, 3506, 3516, 3633, 3715, 3721, 3736, 3744, 3748, 3750, 3756, 3761,
         3781, 3912, 4239, 4347, 4681, 4695, 4697, 4745, 4785, 4799, 4801, 4823,
         4881, 5760, 5901, 5997, 6313, 7405, 8024, 8026, 8028, 8030, 8117, 8125,
         8133, 8181, 8468, 8485, 8487, 8489, 8494, 8527, 11311, 11359, 11687, 11695,
         11703, 11711, 11719, 11727, 11735, 12448, 12539, 43010, 43014, 43019, 43587,
         43696, 43713, 64286, 64297, 64311, 64317, 64319, 64322, 64325, 65141];
    var i, j, start, end;
    for (i = 0; i < singles.length; i++) {
        result[singles[i]] = true;
    }
    var ranges = [[0, 47], [58, 64], [91, 94], [123, 169], [171, 177], [182, 184], [706, 709],
         [722, 735], [741, 747], [751, 879], [888, 889], [894, 901], [1154, 1161],
         [1318, 1328], [1367, 1368], [1370, 1376], [1416, 1487], [1515, 1519], [1523, 1568],
         [1611, 1631], [1642, 1645], [1750, 1764], [1767, 1773], [1789, 1790], [1792, 1807],
         [1840, 1868], [1958, 1968], [1970, 1983], [2027, 2035], [2038, 2041], [2043, 2047],
         [2070, 2073], [2075, 2083], [2085, 2087], [2089, 2307], [2362, 2364], [2366, 2383],
         [2385, 2391], [2402, 2405], [2419, 2424], [2432, 2436], [2445, 2446], [2449, 2450],
         [2483, 2485], [2490, 2492], [2494, 2509], [2511, 2523], [2530, 2533], [2546, 2547],
         [2554, 2564], [2571, 2574], [2577, 2578], [2618, 2648], [2655, 2661], [2672, 2673],
         [2677, 2692], [2746, 2748], [2750, 2767], [2769, 2783], [2786, 2789], [2800, 2820],
         [2829, 2830], [2833, 2834], [2874, 2876], [2878, 2907], [2914, 2917], [2930, 2946],
         [2955, 2957], [2966, 2968], [2976, 2978], [2981, 2983], [2987, 2989], [3002, 3023],
         [3025, 3045], [3059, 3076], [3130, 3132], [3134, 3159], [3162, 3167], [3170, 3173],
         [3184, 3191], [3199, 3204], [3258, 3260], [3262, 3293], [3298, 3301], [3312, 3332],
         [3386, 3388], [3390, 3423], [3426, 3429], [3446, 3449], [3456, 3460], [3479, 3481],
         [3518, 3519], [3527, 3584], [3636, 3647], [3655, 3663], [3674, 3712], [3717, 3718],
         [3723, 3724], [3726, 3731], [3752, 3753], [3764, 3772], [3774, 3775], [3783, 3791],
         [3802, 3803], [3806, 3839], [3841, 3871], [3892, 3903], [3949, 3975], [3980, 4095],
         [4139, 4158], [4170, 4175], [4182, 4185], [4190, 4192], [4194, 4196], [4199, 4205],
         [4209, 4212], [4226, 4237], [4250, 4255], [4294, 4303], [4349, 4351], [4686, 4687],
         [4702, 4703], [4750, 4751], [4790, 4791], [4806, 4807], [4886, 4887], [4955, 4968],
         [4989, 4991], [5008, 5023], [5109, 5120], [5741, 5742], [5787, 5791], [5867, 5869],
         [5873, 5887], [5906, 5919], [5938, 5951], [5970, 5983], [6001, 6015], [6068, 6102],
         [6104, 6107], [6109, 6111], [6122, 6127], [6138, 6159], [6170, 6175], [6264, 6271],
         [6315, 6319], [6390, 6399], [6429, 6469], [6510, 6511], [6517, 6527], [6572, 6592],
         [6600, 6607], [6619, 6655], [6679, 6687], [6741, 6783], [6794, 6799], [6810, 6822],
         [6824, 6916], [6964, 6980], [6988, 6991], [7002, 7042], [7073, 7085], [7098, 7167],
         [7204, 7231], [7242, 7244], [7294, 7400], [7410, 7423], [7616, 7679], [7958, 7959],
         [7966, 7967], [8006, 8007], [8014, 8015], [8062, 8063], [8127, 8129], [8141, 8143],
         [8148, 8149], [8156, 8159], [8173, 8177], [8189, 8303], [8306, 8307], [8314, 8318],
         [8330, 8335], [8341, 8449], [8451, 8454], [8456, 8457], [8470, 8472], [8478, 8483],
         [8506, 8507], [8512, 8516], [8522, 8525], [8586, 9311], [9372, 9449], [9472, 10101],
         [10132, 11263], [11493, 11498], [11503, 11516], [11518, 11519], [11558, 11567],
         [11622, 11630], [11632, 11647], [11671, 11679], [11743, 11822], [11824, 12292],
         [12296, 12320], [12330, 12336], [12342, 12343], [12349, 12352], [12439, 12444],
         [12544, 12548], [12590, 12592], [12687, 12689], [12694, 12703], [12728, 12783],
         [12800, 12831], [12842, 12880], [12896, 12927], [12938, 12976], [12992, 13311],
         [19894, 19967], [40908, 40959], [42125, 42191], [42238, 42239], [42509, 42511],
         [42540, 42559], [42592, 42593], [42607, 42622], [42648, 42655], [42736, 42774],
         [42784, 42785], [42889, 42890], [42893, 43002], [43043, 43055], [43062, 43071],
         [43124, 43137], [43188, 43215], [43226, 43249], [43256, 43258], [43260, 43263],
         [43302, 43311], [43335, 43359], [43389, 43395], [43443, 43470], [43482, 43519],
         [43561, 43583], [43596, 43599], [43610, 43615], [43639, 43641], [43643, 43647],
         [43698, 43700], [43703, 43704], [43710, 43711], [43715, 43738], [43742, 43967],
         [44003, 44015], [44026, 44031], [55204, 55215], [55239, 55242], [55292, 55295],
         [57344, 63743], [64046, 64047], [64110, 64111], [64218, 64255], [64263, 64274],
         [64280, 64284], [64434, 64466], [64830, 64847], [64912, 64913], [64968, 65007],
         [65020, 65135], [65277, 65295], [65306, 65312], [65339, 65344], [65371, 65381],
         [65471, 65473], [65480, 65481], [65488, 65489], [65496, 65497]];
    for (i = 0; i < ranges.length; i++) {
        start = ranges[i][0];
        end = ranges[i][1];
        for (j = start; j <= end; j++) {
            result[j] = true;
        }
    }
    return result;
})();

function splitQuery(query) {
    var result = [];
    var start = -1;
    for (var i = 0; i < query.length; i++) {
        if (splitChars[query.charCodeAt(i)]) {
            if (start !== -1) {
                result.push(query.slice(start, i));
                start = -1;
            }
        } else if (start === -1) {
            start = i;
        }
    }
    if (start !== -1) {
        result.push(query.slice(start));
    }
    return result;
}


