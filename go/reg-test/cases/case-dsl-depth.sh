run_mlr --opprint --from $indir/s.dkvp put '
  $dx = depth($x);
  $dn = depth($nonesuch);
  $da1 = depth([1,2,3]);
  $da2 = depth([1,[4,5,6],3]);
  $da3 = depth([1,{"s":4,"t":[7,8,9],"u":6},3]);
  $dm1 = depth({"s":1,"t":2,"u":3});
  $dm2 = depth({"s":1,"t":[4,5,6],"u":3});
  $dm3 = depth({"s":1,"t":[4,$*,6],"u":3});
'
