package validation

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testdata = `
{
  "tag": "陸上自衛隊",
  "geotags": [
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8282/7831339680_265efb968e.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7124/7831342090_9383001e63.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8435/7831344540_45f9d97279.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7271/7831347324_be3bfc6e59.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8286/7831350858_071716715a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8429/7831354006_0e9c3e7e3e.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8448/7831357238_bf1a2da3f9.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8296/7831359418_025032f36a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7127/7831360968_c9023e994c.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8431/7831363944_a7a3321109.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8443/7831366386_f4d9b29e22.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8300/7831368906_24b4eb700a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8425/7831371914_00575107e2.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8290/7831374036_eddf6cc182.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7139/7831375802_d99980af5d.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8024/7831377610_19653619ff.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8424/7831380546_41e17cbe9d.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7126/7831382416_792f1dfb56.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8307/7831384168_2b1704aa88.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8286/7831386130_afe1c041ee.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8298/7831388268_e8b03c2486.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7274/7831389862_19bedc2c5a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8296/7831392154_ef82e91dca.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8287/7831395164_b565e23725.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7253/7831398070_2286a398d2.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7123/7831400576_a9bda100b1.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8438/7831404458_9ceee2f4cc.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7123/7831407562_cddb939312.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8422/7831409974_d9231309a2.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8287/7831411816_72c88801da.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8439/7831413796_a1fe26caa4.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8282/7831415662_e1f2323a49.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8306/7831417576_97d55150e3.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7273/7831419838_e49a7fe605.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7255/7831421436_08b6404ac0.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7255/7831423508_b3d74f823e.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7127/7831426138_12408a3ae6.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8293/7831429114_39db0ca01e.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8305/7831431188_da322837d6.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8283/7831432806_e18a8a173d.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8432/7831435650_e47893e97b.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8290/7831437584_33aec12883.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8303/7831439654_001883f2da.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8303/7831442114_36299015b2.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8307/7831444880_2d70af4bb0.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7107/7831448264_6a772ab950.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7130/7831450392_b49f38307a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7250/7831451868_1664a44077.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8292/7831453880_345b45122b.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8286/7831458552_06f71c597a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8306/7831461014_78f3e47a3e.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7266/7831463136_473529869a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8435/7831464680_c53eeebd24.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8442/7831467106_7b488136b0.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8422/7831468784_b78d7ec313.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8291/7831470684_9e3a7dfffe.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7246/7831473190_e74ec43663.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8296/7831474886_9e5e656665.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8290/7831477270_6ea5c42f69.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8423/7831478920_bcd7662610.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8304/7831482574_f035ff8b5f.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8437/7831485036_99448a244a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7121/7831486338_6cbe3c3ba3.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7124/7831487788_c2ba301b4d.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8286/7831489506_81be887c49.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7267/7831492566_e850a0a1e3.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8281/7831494094_fe5edebd24.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8438/7831497236_bf1ae380ed.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8429/7831500464_4abb69c51b.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7263/7831503432_7eaee37aeb.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8296/7831506334_1cdba14769.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7260/7831508842_6d3b724e61.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8292/7831512352_4b6db0d74d.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8289/7831514666_6d4bb04f05.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8291/7831517582_5cc77a9cf8.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7257/7831523602_4794092dbd.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8435/7831520118_3b29382947.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8435/7831521910_7c5831432b.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7257/7831525850_d460e686de.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8422/7831528052_8269fc40c7.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8427/7831530904_a0fba04bbd.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7126/7831534014_15fcb732c4.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8291/7831536646_e8cf0afc28.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8424/7831538880_241c8e3f3a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7279/7831540748_f5b67c7ef5.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7252/7831543452_6bf06c0a9b.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8446/7831546130_d1f0e33a44.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7124/7831549648_b41f72b259.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8426/7831551708_5658e02a87.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8447/7831552796_0952a8bf92.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7255/7831555332_af0e55406f.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7255/7831558228_dc47803dae.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8421/7862512996_4fb176b1f5.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8443/7862514270_af856108d3.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7262/7862515126_d9e0c48c3a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8287/7862516378_1485ddf00a.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8301/7862517760_1e8120a5df.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8301/7862519068_32fcf27b92.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm9.static.flickr.com/8284/7862520418_f33289cc04.jpg"
    },
    {
      "lat": 35.285423,
      "lon": 138.84761,
      "date": "2012-08-21UTC",
      "url": "https://farm8.static.flickr.com/7127/7862521276_1c7cfcff12.jpg"
    }
  ]
}`

func Test_validate2022(t *testing.T) {
	uri, _ := url.ParseRequestURI("http://localhost:8080/program?tag=陸上自衛隊")
	err := Validate2022(uri, []byte(testdata))
	assert.NoError(t, err)
}
