		</div><!-- content -->	
	</section>
</div><!-- wrapper -->

<?php apply_hooks('tb_footer'); ?>

<div class="footer-toolbar" id="footer-toolbar">
	<div class="back-to-top" id="back-to-top" title="回到顶端">
		<i class="fa fa-arrow-circle-up"></i>
	</div>
	<div style="display: none;" class="reading-mode no-sel" id="reading-mode" title="阅读模式">
		<i class="fa fa-plus-circle"></i>
	</div>
</div>
<div class="img-view" id="img-view"><img /><div class="tip"></div></div>
<div style="display: none;">
<?php if(!$tbmain->is_ssl) echo '<script src="http://js.users.51.la/17768957.js"></script>'; ?>
	<script>
		(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
		(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
		m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
		})(window,document,'script','//www.google-analytics.com/analytics.js','ga');

		ga('create', 'UA-65174773-1', 'auto');
		ga('send', 'pageview');
	</script>
    <?php if(!$tbmain->is_ssl) : ?>
    <script>
        (function(){
            var bp = document.createElement('script');
            bp.src = '//push.zhanzhang.baidu.com/push.js';
            var s = document.getElementsByTagName("script")[0];
            s.parentNode.insertBefore(bp, s);
        })();
    </script>
<?php endif; ?>
</div>
<script src="/theme/scripts/footer.js"></script>
</body>
</html>
<!-- 执行时间: <?php echo $execution_time; ?>s -->
<?php

