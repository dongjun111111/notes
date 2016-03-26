###Geohash的认知
Geohash的最简单的解释就是：将一个经纬度信息，转换成一个可以排序，可以比较的字符串编码.
###Geohash的特点
- 首先，geohash用一个字符串表示经度和纬度两个坐标。某些情况下无法在两列上同时应用索引 （例如MySQL 4之前的版本，Google App Engine的数据层等），利用geohash，只需在一列上应用索引即可。
- 其次，geohash表示的并不是一个点，而是一个矩形区域。比如编码wx4g0ec19，它表示的是一个矩形区域。 使用者可以发布地址编码，既能表明自己位于北海公园附近，又不至于暴露自己的精确坐标，有助于隐私保护。
- 第三，编码的前缀可以表示更大的区域。例如wx4g0ec1，它的前缀wx4g0e表示包含编码wx4g0ec1在内的更大范围。 这个特性可以用于附近地点搜索。首先根据用户当前坐标计算geohash（例如wx4g0ec1）然后取其前缀进行查询 （SELECT * FROM place WHERE geohash LIKE 'wx4g0e%'），即可查询附近的所有地点。

当然，Geohash比直接用经纬度的高效很多。

###Geohash的原理
首先将纬度范围(-90, 90)平分成两个区间(-90,0)、(0, 90)，如果目标纬度位于前一个区间，则编码为0，否则编码为1。

由于39.92324属于(0, 90)，所以取编码为1。

然后再将(0, 90)分成 (0, 45), (45, 90)两个区间，而39.92324位于(0, 45)，所以编码为0。

以此类推，直到精度符合要求为止，得到纬度编码为1011 1000 1100 0111 1001。

经度也用同样的算法，对(-180, 180)依次细分，得到116.3906的编码为1101 0010 1100 0100 0100。

接下来将经度和纬度的编码合并，奇数位是纬度，偶数位是经度，得到编码 11100 11101 00100 01111 00000 01101 01011 00001。

最后，用0-9、b-z（去掉a, i, l, o）这32个字母进行base32编码，得到(39.92324, 116.3906)的编码为wx4g0ec1。
<table>
	<thead>
		<tr><th>十进制</th><th>0</th><th>1</th><th>2</th><th>3</th><th>4</th><th>5</th><th>6</th><th>7</th><th>8</th><th>9</th><th>10</th><th>11</th><th>12</th><th>13</th><th>14</th><th>15</th></tr>
	</thead>
	<tbody>
		<tr><td>base32</td><td>0</td><td>1</td><td>2</td><td>3</td><td>4</td><td>5</td><td>6</td><td>7</td><td>8</td><td>9</td><td>b</td><td>c</td><td>d</td><td>e</td><td>f</td><td>g</td></tr>
		<tr><td>十进制</td><td>16</td><td>17</td><td>18</td><td>19</td><td>20</td><td>21</td><td>22</td><td>23</td><td>24</td><td>25</td><td>26</td><td>27</td><td>28</td><td>29</td><td>30</td><td>31</td></tr>
		<tr><td>base32</td><td>h</td><td>j</td><td>k</td><td>m</td><td>n</td><td>p</td><td>q</td><td>r</td><td>s</td><td>t</td><td>u</td><td>v</td><td>w</td><td>x</td><td>y</td><td>z</td></tr>
	
	</tbody>
</table>
解码算法与编码算法相反，先进行base32解码，然后分离出经纬度，最后根据二进制编码对经纬度范围进行细分即可，这里不再赘述。

实现代码(php版)：
<pre>
/**
 * Geohash generation class
 * http://blog.dixo.net/downloads/
 *
 * This file copyright (C) 2008 Paul Dixon (paul@elphin.com)
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
 */



/**
* Encode and decode geohashes
*
*/
class Geohash
{
    private $coding="0123456789bcdefghjkmnpqrstuvwxyz";
    private $codingMap=array();
    
    public function Geohash()
    {
        //build map from encoding char to 0 padded bitfield
        for($i=0; $i<32; $i++)
        {
            $this->codingMap[substr($this->coding,$i,1)]=str_pad(decbin($i), 5, "0", STR_PAD_LEFT);
        }
        
    }
    
    /**
    * Decode a geohash and return an array with decimal lat,long in it
    */
    public function decode($hash)
    {
        //decode hash into binary string
        $binary="";
        $hl=strlen($hash);
        for($i=0; $i<$hl; $i++)
        {
            $binary.=$this->codingMap[substr($hash,$i,1)];
        }
        
        //split the binary into lat and log binary strings
        $bl=strlen($binary);
        $blat="";
        $blong="";
        for ($i=0; $i<$bl; $i++)
        {
            if ($i%2)
                $blat=$blat.substr($binary,$i,1);
            else
                $blong=$blong.substr($binary,$i,1);
            
        }
        
        //now concert to decimal
        $lat=$this->binDecode($blat,-90,90);
        $long=$this->binDecode($blong,-180,180);
        
        //figure out how precise the bit count makes this calculation
        $latErr=$this->calcError(strlen($blat),-90,90);
        $longErr=$this->calcError(strlen($blong),-180,180);
                
        //how many decimal places should we use? There's a little art to
        //this to ensure I get the same roundings as geohash.org
        $latPlaces=max(1, -round(log10($latErr))) - 1;
        $longPlaces=max(1, -round(log10($longErr))) - 1;
        
        //round it
        $lat=round($lat, $latPlaces);
        $long=round($long, $longPlaces);
        
        return array($lat,$long);
    }

    
    /**
    * Encode a hash from given lat and long
    */
    public function encode($lat,$long)
    {
        //how many bits does latitude need?    
        $plat=$this->precision($lat);
        $latbits=1;
        $err=45;
        while($err>$plat)
        {
            $latbits++;
            $err/=2;
        }
        
        //how many bits does longitude need?
        $plong=$this->precision($long);
        $longbits=1;
        $err=90;
        while($err>$plong)
        {
            $longbits++;
            $err/=2;
        }
        
        //bit counts need to be equal
        $bits=max($latbits,$longbits);
        
        //as the hash create bits in groups of 5, lets not
        //waste any bits - lets bulk it up to a multiple of 5
        //and favour the longitude for any odd bits
        $longbits=$bits;
        $latbits=$bits;
        $addlong=1;
        while (($longbits+$latbits)%5 != 0)
        {
            $longbits+=$addlong;
            $latbits+=!$addlong;
            $addlong=!$addlong;
        }
        
        
        //encode each as binary string
        $blat=$this->binEncode($lat,-90,90, $latbits);
        $blong=$this->binEncode($long,-180,180,$longbits);
        
        //merge lat and long together
        $binary="";
        $uselong=1;
        while (strlen($blat)+strlen($blong))
        {
            if ($uselong)
            {
                $binary=$binary.substr($blong,0,1);
                $blong=substr($blong,1);
            }
            else
            {
                $binary=$binary.substr($blat,0,1);
                $blat=substr($blat,1);
            }
            $uselong=!$uselong;
        }
        
        //convert binary string to hash
        $hash="";
        for ($i=0; $i<strlen($binary); $i+=5)
        {
            $n=bindec(substr($binary,$i,5));
            $hash=$hash.$this->coding[$n];
        }
        
        
        return $hash;
    }
    
    /**
    * What's the maximum error for $bits bits covering a range $min to $max
    */
    private function calcError($bits,$min,$max)
    {
        $err=($max-$min)/2;
        while ($bits--)
            $err/=2;
        return $err;
    }
    
    /*
    * returns precision of number
    * precision of 42 is 0.5
    * precision of 42.4 is 0.05
    * precision of 42.41 is 0.005 etc
    */
    private function precision($number)
    {
        $precision=0;
        $pt=strpos($number,'.');
        if ($pt!==false)
        {
            $precision=-(strlen($number)-$pt-1);
        }
        
        return pow(10,$precision)/2;
    }
    
    
    /**
    * create binary encoding of number as detailed in http://en.wikipedia.org/wiki/Geohash#Example
    * removing the tail recursion is left an exercise for the reader
    */
    private function binEncode($number, $min, $max, $bitcount)
    {
        if ($bitcount==0)
            return "";
        
        #echo "$bitcount: $min $max<br>";
            
        //this is our mid point - we will produce a bit to say
        //whether $number is above or below this mid point
        $mid=($min+$max)/2;
        if ($number>$mid)
            return "1".$this->binEncode($number, $mid, $max,$bitcount-1);
        else
            return "0".$this->binEncode($number, $min, $mid,$bitcount-1);
    }
    

    /**
    * decodes binary encoding of number as detailed in http://en.wikipedia.org/wiki/Geohash#Example
    * removing the tail recursion is left an exercise for the reader
    */
    private function binDecode($binary, $min, $max)
    {
        $mid=($min+$max)/2;
        
        if (strlen($binary)==0)
            return $mid;
            
        $bit=substr($binary,0,1);
        $binary=substr($binary,1);
        
        if ($bit==1)
            return $this->binDecode($binary, $mid, $max);
        else
            return $this->binDecode($binary, $min, $mid);
    }
}
?>
</pre>


