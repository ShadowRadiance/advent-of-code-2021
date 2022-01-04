# frozen_string_literal: true

require "./dirac_dice/game"
require "./dirac_dice/deterministic_die"

RSpec.describe DiracDice::Game do
  let(:die) { DeterministicDie.new(100) }
  let(:game) { described_class.new(p1_start, p2_start, die: die) }
  let(:p1_start) { 4 }
  let(:p2_start) { 8 }

  before { game.play }

  it "determines the correct score" do
    expect(game.loser.score).to eq(745)
    expect(die.times_rolled).to eq(993)
    expect(game.score).to eq(739785)
  end
end