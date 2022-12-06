# frozen_string_literal: true

require "./dirac_dice/game"
require "./dirac_dice/game2"
require "./dirac_dice/deterministic_die"

RSpec.describe "Dice Game" do
  let(:p1_start) { 4 }
  let(:p2_start) { 8 }

  context "in the deterministic die scenario" do
    let(:die) { DeterministicDie.new(100) }
    let(:game) { DiracDice::Game.new(p1_start, p2_start, die: die) }
      
    before { game.play }

    it "determines the correct score" do
      expect(game.loser.score).to eq(745)
      expect(die.times_rolled).to eq(993)
      expect(game.score).to eq(739785)
    end
  end

  context "in a multiversal scenario" do
    let(:game) { DiracDice::Game2.new(p1_start, p2_start) }
    before { game.play }

    it "determines the correct score" do
      expect(game.player_one_wins).to eq(444356092776315)
      expect(game.player_two_wins).to eq(341960390180808)
    end
  end

end