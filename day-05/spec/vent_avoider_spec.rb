require 'byebug'
require './vent/avoider'

RSpec.describe Vent::Avoider do
  subject { described_class.new(data) }

  let(:data) {
    <<~DATA
    0,9 -> 5,9
    8,0 -> 0,8
    9,4 -> 3,4
    2,2 -> 2,1
    7,0 -> 7,4
    6,4 -> 2,0
    0,9 -> 2,9
    3,4 -> 1,4
    0,0 -> 8,8
    5,5 -> 8,2
    DATA
  }

  it "counts the overlapping grid locations" do
    expect(subject.overlapping_points.size).to eq(5)
  end
end