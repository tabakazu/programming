# 参考: https://matz.rubyist.net/20071122.html#p01

# 開け閉めできるドアクラス
class Door
  def close; puts 'ドアを閉めました' end
  def closed?; false end
end

# 開け閉めできるカーテンクラス
class Curtain
  def close; puts 'カーテンを閉めました' end
  def closed?; false end
end

# 開け閉めできる DbConnect クラス
class DbConnect
  def close; puts 'DB をクローズしました' end
  def closed?; false end
end

def safe_close(x)
  x.close unless x.closed? 
end

# ドアを作って、閉める
door = Door.new
safe_close(door)

# カーテンを作って、閉める
curtain = Curtain.new
safe_close(curtain)

# DbConnect を作って、閉める
db = DbConnect.new
safe_close(db)