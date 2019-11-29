class User < ActiveRecord::Base
  include Hashid::Rails
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable,
         :confirmable, :omniauthable

  # Association
  belongs_to :account
  belongs_to :organization, optional: true
  has_many :posts, dependent: :destroy

  # Accepts nested attributes for
  accepts_nested_attributes_for :account

  # Validation
  validates :name, presence: true, uniqueness: true

  # Scope
  scope :active, -> { where(active: true) }
  scope :matches_email, -> (email) { where(arel_table[:email].matches("%#{email}%")) }
  scope :recent, -> { order(created_at: :desc) }

  # Delegate
  delegate :name, :display_name, :gender, :description, :avatar,
           to: :account, prefix: true, allow_nil: true

  # Instance Method
  def recent_active_posts
    posts.recently_active
  end

  # Class Method
  class << self
    def emails
      pluck(:email)
    end
  end
end
