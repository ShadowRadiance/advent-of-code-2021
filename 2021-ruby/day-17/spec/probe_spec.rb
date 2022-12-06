# frozen_string_literal: true
require './probe'
require 'matrix'

RSpec.describe Probe do
  context "with no parameters" do
    it "has not moved" do
      expect(subject.location).to eq(Vector[0,0])
    end

    it "is not moving" do
      expect(subject.velocity).to eq(Vector[0,0])
    end

    context "after one step" do
      before { subject.step }

      it "has not moved yet" do
        expect(subject.location).to eq(Vector[0,0])
      end

      it "is starting to fall" do
        expect(subject.velocity).to eq(Vector[0,-1])
      end

      context "after another step" do
        before { subject.step }

        it "has started falling" do
          expect(subject.location).to eq(Vector[0,-1])
        end

        it "is falling faster" do
          expect(subject.velocity).to eq(Vector[0,-2])
        end
      end
    end
  end

  context "with an initial velocity of 7,2" do
    subject { Probe.new(7,2) }
    context "before any steps" do
      it "has not moved and has velocity 7,2" do
        expect(subject.location).to eq(Vector[0,0])
        expect(subject.velocity).to eq(Vector[7,2])
      end
    end
    context "after 7 steps" do
      before { 7.times { subject.step } }
      it "has moved to 28,-7" do
        expect(subject.location).to eq(Vector[28,-7])
      end
    end
  end
end
