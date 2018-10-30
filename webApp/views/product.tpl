<!DOCTYPE html>
<html>
   <head>
      <title>产品录入</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <!-- 引入 Bootstrap -->
      <link href="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
 
      <!-- HTML5 Shiv 和 Respond.js 用于让 IE8 支持 HTML5元素和媒体查询 -->
      <!-- 注意： 如果通过 file://  引入 Respond.js 文件，则该文件无法起效果 -->
      <!--[if lt IE 9]>
         <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
         <script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
      <![endif]-->
   </head>
   <body>
   <div class = "container">
      <form role="form" method="post" action ="/product">
		  <div class="form-group">
			<label for="name">产品名称</label>
			<input type="text" class="form-control" name ="Name" id="name" placeholder="产品名称">
		  </div>
		   <div class="form-group">
			<label for="name">进货价</label>
			<input type="text" class="form-control" name="Order_price" id="order_price" placeholder="进货价">
		  </div>
		  <div class="form-group">
			<label for="name">卖价</label>
			<input type="text" class="form-control" name="Origin_price" id="origin_price" placeholder="卖价">
		  </div>
		  <div class="form-group">
			<label for="name">折扣</label>
			<input type="text" class="form-control" name ="Discount" id="discount" placeholder="折扣">
		  </div>

		   <div class="form-group">
			<label for="name">起卖数量</label>
			<input type="text" id="min_num" class="form-control" name="Min_num" value="1" >
		  </div>
		  <div class="form-group">
			<label for="name">单位</label>
			<input type="text" id="unit" class="form-control" name="Unit"    value="支">
		  </div>
		  <button type="submit" class="btn btn-default">提交</button>
	</form>
	</div>
      <!-- jQuery (Bootstrap 的 JavaScript 插件需要引入 jQuery) -->
      <script src="https://cdn.staticfile.org/jquery/2.1.1/jquery.min.js"></script>
      <!-- 包括所有已编译的插件 -->
      <script src="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
   </body>
</html>