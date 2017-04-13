#!/usr/bin/env php
<?php
require './unu.php';

$unu = (new UNURequest())
	->setURL('https://xtom.com/');

var_dump($unu->submit());